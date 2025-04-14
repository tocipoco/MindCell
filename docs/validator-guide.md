# Validator Setup Guide

## Overview

Validators are critical to MindCell's consensus mechanism. They propose blocks, validate transactions, and participate in governance.

## Hardware Requirements

### Minimum Specifications
- CPU: 8 cores / 16 threads @ 2.5GHz
- RAM: 32GB DDR4
- Storage: 1TB NVMe SSD
- Network: 100Mbps symmetric with static IP
- OS: Ubuntu 22.04 LTS or similar

### Recommended for Production
- CPU: 16 cores / 32 threads @ 3.5GHz+
- RAM: 64GB DDR4
- Storage: 2TB NVMe SSD (RAID 1 or RAID 10)
- Network: 1Gbps symmetric with redundant connections
- OS: Ubuntu 22.04 LTS (hardened)

## Initial Setup

### Install Dependencies

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install build tools
sudo apt install -y build-essential git curl wget jq

# Install Go 1.21+
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Set up Go environment
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Build MindCell

```bash
# Clone repository
git clone https://github.com/tocipoco/MindCell.git
cd MindCell

# Build binary
make install

# Verify installation
mindcelld version
```

## Node Initialization

### Generate Node Key and Config

```bash
# Initialize node
mindcelld init my-validator --chain-id mindcell-1

# This creates:
# ~/.mindcell/config/config.toml
# ~/.mindcell/config/app.toml
# ~/.mindcell/config/genesis.json
# ~/.mindcell/config/node_key.json
# ~/.mindcell/config/priv_validator_key.json
```

### Download Genesis File

```bash
# Download official genesis
curl https://raw.githubusercontent.com/mindcell-network/networks/main/mainnet/genesis.json \
  > ~/.mindcell/config/genesis.json

# Verify genesis hash
sha256sum ~/.mindcell/config/genesis.json
# Should match: abc123...official-hash
```

### Configure Seeds and Persistent Peers

```bash
# Edit config.toml
nano ~/.mindcell/config/config.toml

# Set persistent peers
persistent_peers = "node1@ip1:26656,node2@ip2:26656,node3@ip3:26656"

# Set seeds
seeds = "seed1@ip:26656,seed2@ip:26656"

# Set external address (your public IP)
external_address = "tcp://your.public.ip:26656"
```

## Security Configuration

### Firewall Setup

```bash
# Allow SSH
sudo ufw allow 22/tcp

# Allow P2P
sudo ufw allow 26656/tcp

# Allow RPC (only from trusted IPs)
sudo ufw allow from trusted_ip to any port 26657

# Enable firewall
sudo ufw enable
```

### Key Management

```bash
# Backup validator key (CRITICAL!)
cp ~/.mindcell/config/priv_validator_key.json ~/validator_key_backup.json

# Set proper permissions
chmod 600 ~/.mindcell/config/priv_validator_key.json

# Consider using a KMS for production
# Reference: https://github.com/iqlusioninc/tmkms
```

### Sentry Node Architecture

Set up sentry nodes to protect validator from DDoS:

```
[Public Internet]
       ↓
  [Sentry Node 1] ← [Sentry Node 2] ← [Sentry Node 3]
       ↓                  ↓                  ↓
            [Validator Node (Private)]
```

## Create Validator

### Add Account

```bash
# Create new account
mindcelld keys add validator

# Or recover existing
mindcelld keys add validator --recover

# Get address
mindcelld keys show validator -a
```

### Get Tokens

```bash
# For testnet, use faucet
curl -X POST https://faucet.mindcell.network/request \
  -d '{"address":"cosmos1..."}'

# For mainnet, acquire MCELL from exchanges
```

### Submit Create-Validator Transaction

```bash
mindcelld tx staking create-validator \
  --amount=100000mcell \
  --pubkey=$(mindcelld tendermint show-validator) \
  --moniker="My Validator" \
  --identity="keybase_identity" \
  --website="https://myvalidator.com" \
  --security-contact="security@myvalidator.com" \
  --details="Dedicated MindCell validator" \
  --chain-id=mindcell-1 \
  --commission-rate=0.10 \
  --commission-max-rate=0.20 \
  --commission-max-change-rate=0.01 \
  --min-self-delegation=1 \
  --gas=auto \
  --gas-adjustment=1.5 \
  --from=validator \
  --fees=1000mcell
```

## Running as Service

### Create Systemd Service

```bash
sudo tee /etc/systemd/system/mindcelld.service > /dev/null <<EOF
[Unit]
Description=MindCell Validator Node
After=network-online.target

[Service]
User=$USER
ExecStart=$(which mindcelld) start
Restart=always
RestartSec=3
LimitNOFILE=65535
Environment="DAEMON_HOME=$HOME/.mindcell"
Environment="DAEMON_NAME=mindcelld"

[Install]
WantedBy=multi-user.target
EOF

# Enable and start service
sudo systemctl daemon-reload
sudo systemctl enable mindcelld
sudo systemctl start mindcelld

# Check status
sudo systemctl status mindcelld

# View logs
journalctl -u mindcelld -f
```

## Monitoring

### Prometheus Metrics

Enable Prometheus in `app.toml`:

```toml
[telemetry]
enabled = true
prometheus-retention-time = 60
```

Access metrics at: `http://localhost:26660/metrics`

### Key Metrics to Monitor

```yaml
# Validator status
tendermint_consensus_validator_power
tendermint_consensus_validators

# Block production
tendermint_consensus_height
tendermint_consensus_rounds

# Network health
tendermint_p2p_peers
tendermint_consensus_missing_validators

# Performance
tendermint_mempool_size
tendermint_consensus_block_interval_seconds
```

### Alerting Setup

```yaml
# Prometheus alert rules
groups:
  - name: validator_alerts
    rules:
      - alert: ValidatorDown
        expr: up{job="mindcell-validator"} == 0
        for: 5m
        annotations:
          summary: "Validator is down"
          
      - alert: MissingBlocks
        expr: increase(tendermint_consensus_missing_validators[10m]) > 5
        annotations:
          summary: "Validator missing blocks"
          
      - alert: LowPeerCount
        expr: tendermint_p2p_peers < 5
        for: 10m
        annotations:
          summary: "Low peer count"
```

## Maintenance

### Upgrade Procedures

```bash
# Stop node
sudo systemctl stop mindcelld

# Backup state
cp -r ~/.mindcell/data ~/.mindcell/data_backup

# Download new binary
wget https://github.com/tocipoco/MindCell/releases/download/v2.0.0/mindcelld
chmod +x mindcelld
sudo mv mindcelld /usr/local/bin/

# Restart node
sudo systemctl start mindcelld

# Monitor sync
mindcelld status 2>&1 | jq .SyncInfo
```

### State Sync (Fast Sync)

For new validators to catch up quickly:

```bash
# Get trusted height and hash
LATEST_HEIGHT=$(curl -s https://rpc.mindcell.network/block | jq -r .result.block.header.height)
TRUST_HEIGHT=$((LATEST_HEIGHT - 1000))
TRUST_HASH=$(curl -s "https://rpc.mindcell.network/block?height=$TRUST_HEIGHT" | jq -r .result.block_id.hash)

# Configure state sync
nano ~/.mindcell/config/config.toml

[statesync]
enable = true
rpc_servers = "https://rpc1.mindcell.network:443,https://rpc2.mindcell.network:443"
trust_height = $TRUST_HEIGHT
trust_hash = "$TRUST_HASH"
trust_period = "168h"  # 1 week
```

### Snapshot Restore

```bash
# Download latest snapshot
wget https://snapshots.mindcell.network/latest.tar.gz

# Stop node
sudo systemctl stop mindcelld

# Remove old data
rm -rf ~/.mindcell/data

# Extract snapshot
tar -xzf latest.tar.gz -C ~/.mindcell/

# Restart node
sudo systemctl start mindcelld
```

## Validator Operations

### Unjail Validator

If validator gets jailed due to downtime:

```bash
mindcelld tx slashing unjail \
  --from=validator \
  --chain-id=mindcell-1 \
  --fees=500mcell
```

### Edit Validator Info

```bash
mindcelld tx staking edit-validator \
  --website="https://newwebsite.com" \
  --details="Updated description" \
  --commission-rate=0.08 \
  --from=validator \
  --chain-id=mindcell-1
```

### Delegate More Tokens

```bash
# Self-delegation
mindcelld tx staking delegate \
  $(mindcelld keys show validator --bech val -a) \
  50000mcell \
  --from=validator \
  --chain-id=mindcell-1
```

## Governance Participation

### View Active Proposals

```bash
mindcelld query gov proposals --status=voting_period
```

### Vote on Proposal

```bash
# Vote: yes, no, abstain, no_with_veto
mindcelld tx gov vote 1 yes \
  --from=validator \
  --chain-id=mindcell-1 \
  --fees=100mcell
```

### Submit Proposal

```bash
# Parameter change proposal
mindcelld tx gov submit-proposal param-change proposal.json \
  --from=validator \
  --deposit=10000mcell \
  --chain-id=mindcell-1
```

## Troubleshooting

### Node Not Syncing

**Check peers:**
```bash
curl localhost:26657/net_info | jq .result.peers
```

**Solutions:**
- Add more persistent peers
- Check firewall rules
- Verify network connectivity

### High Memory Usage

**Check resource usage:**
```bash
htop
df -h
```

**Solutions:**
- Enable pruning
- Reduce cache sizes
- Add more RAM

### Missed Blocks

**Check signing info:**
```bash
mindcelld query slashing signing-info \
  $(mindcelld tendermint show-validator)
```

**Common causes:**
- High latency
- Resource constraints
- Network issues

**Solutions:**
- Optimize hardware
- Improve network connection
- Check consensus parameters

## Best Practices

1. **Monitor 24/7**: Set up automated alerts
2. **Backup regularly**: Keys, config, and state
3. **Update promptly**: Stay on latest stable version
4. **Participate in governance**: Vote on proposals
5. **Maintain uptime**: 99.9%+ recommended
6. **Secure infrastructure**: Use HSM, firewalls, monitoring
7. **Community engagement**: Join validator chat
8. **Document procedures**: Internal runbooks
9. **Test upgrades**: On testnet first
10. **Plan redundancy**: Have backup systems ready

## Resources

- Validator chat: https://discord.gg/mindcell-validators
- Status page: https://status.mindcell.network
- Block explorer: https://explorer.mindcell.network
- Upgrade coordinator: https://github.com/mindcell-network/upgrades
