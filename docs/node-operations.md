# Node Operations Guide

## Becoming a Node Operator

### Requirements
- Minimum stake: 1,000 MCELL
- Hardware: 8 CPU cores, 16GB RAM, 500GB SSD
- Network: Stable connection with low latency
- Uptime: 99%+ recommended

### Registration

```bash
mindcelld tx shardallocator register-node \
  --endpoint="https://mynode.example.com" \
  --stake=10000mcell \
  --max-shards=50 \
  --from=mykey
```

### Shard Management

Shards are automatically assigned based on:
- Available capacity
- Reputation score
- Network topology

### Monitoring

Monitor your node performance:

```bash
mindcelld query shardallocator node-info <address>
```

### Rewards

Claim accumulated rewards:

```bash
mindcelld tx reward claim-reward --from=mykey
```

### Slashing Conditions

Your stake can be slashed for:
- Timeout: 5% slash
- Incorrect proof: 20% slash
- Prolonged downtime: 10% slash

### Best Practices

1. Maintain high uptime
2. Monitor system resources
3. Keep software updated
4. Backup private keys securely
5. Join the operator community

