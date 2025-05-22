# Network Upgrade Guide

## Overview

Network upgrades are coordinated through on-chain governance proposals.

## Upgrade Process

### 1. Preparation Phase

**Core Team:**
- Develop and test new version
- Create detailed upgrade proposal
- Publish upgrade documentation

**Validators:**
- Review proposed changes
- Test on private testnet
- Prepare upgrade procedures

### 2. Governance Proposal

```bash
# Submit software upgrade proposal
mindcelld tx gov submit-proposal software-upgrade v2.0.0 \
  --title="Upgrade to v2.0.0" \
  --description="Major upgrade with FHE support" \
  --upgrade-height=1000000 \
  --upgrade-info='{"binaries": {"linux/amd64": "https://..."}}' \
  --deposit=50000mcell \
  --from=proposer
```

### 3. Voting Period

Validators and token holders vote:

```bash
# Vote yes
mindcelld tx gov vote 1 yes --from=validator

# Check voting progress
mindcelld query gov tally 1
```

### 4. Upgrade Execution

When block height reaches upgrade height:

**Automatic Halt:**
- Chain automatically halts at upgrade height
- Validators see: "UPGRADE NEEDED: v2.0.0"

**Manual Steps:**
```bash
# Stop node
sudo systemctl stop mindcelld

# Backup state
cp -r ~/.mindcell/data ~/.mindcell/data_backup_pre_v2

# Download new binary
wget https://github.com/tocipoco/MindCell/releases/download/v2.0.0/mindcelld
chmod +x mindcelld
sudo mv mindcelld $(which mindcelld)

# Verify version
mindcelld version
# Should show: v2.0.0

# Restart node
sudo systemctl start mindcelld

# Monitor logs
journalctl -u mindcelld -f
```

**Cosmovisor (Automatic):**
```bash
# Cosmovisor handles binary switching automatically
# Node upgrades and restarts without manual intervention
```

## Upgrade Types

### Minor Upgrades (1.0.x -> 1.1.0)

- Bug fixes and optimizations
- No breaking changes
- Simple binary replacement
- No data migration needed

### Major Upgrades (1.x -> 2.0.0)

- Breaking API changes
- New features
- May require data migration
- Coordinated network halt

## Rollback Procedures

If upgrade fails:

```bash
# Stop node
sudo systemctl stop mindcelld

# Restore backup
rm -rf ~/.mindcell/data
cp -r ~/.mindcell/data_backup_pre_v2 ~/.mindcell/data

# Restore previous binary
sudo mv mindcelld.backup $(which mindcelld)

# Restart
sudo systemctl start mindcelld
```

## Testing Upgrades

### On Testnet

1. Deploy upgrade to testnet first
2. Monitor for 48+ hours
3. Verify all modules function correctly
4. Measure performance impact

### Local Simulation

```bash
# Start local network
./scripts/init-testnet.sh

# Simulate upgrade at block 100
mindcelld tx gov submit-proposal software-upgrade test-upgrade \
  --upgrade-height=100 \
  --from=test1

# Vote and wait for execution
```

## Communication

### Before Upgrade

- Announcement 2 weeks in advance
- Technical specification published
- Validator briefing calls
- Documentation updates

### During Upgrade

- Real-time coordination channel
- Status updates every 15 minutes
- Emergency contact information

### After Upgrade

- Post-mortem analysis
- Performance metrics review
- Issue tracking and resolution

## Emergency Procedures

### Failed Upgrade

If network cannot restart:

1. Emergency coordination call
2. Identify root cause
3. Prepare hotfix
4. Coordinated manual intervention

### State Corruption

If state becomes corrupted:

1. Halt network immediately
2. Validators restore from backup
3. Identify corruption source
4. Apply fix and restart

## Resources

- Upgrade announcements: https://forum.mindcell.network/upgrades
- Coordinator: https://github.com/mindcell-network/upgrades
- Validator chat: Discord #validators
