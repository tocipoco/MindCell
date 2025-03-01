# Troubleshooting Guide

## Common Issues

### Node Won't Start

**Symptom**: Node fails to start or crashes immediately

**Solutions**:
1. Check configuration files for syntax errors
2. Ensure ports are not already in use
3. Verify disk space availability
4. Check system logs for detailed errors

### Sync Issues

**Symptom**: Node not syncing with network

**Solutions**:
1. Verify peer connections
2. Check network connectivity
3. Reset state if corrupted
4. Use state sync for faster sync

### Transaction Failures

**Symptom**: Transactions fail to execute

**Solutions**:
1. Check account balance for sufficient funds
2. Verify gas settings
3. Ensure nonce/sequence is correct
4. Check transaction format

### High Memory Usage

**Symptom**: Node consuming excessive memory

**Solutions**:
1. Enable pruning in configuration
2. Reduce cache sizes
3. Upgrade hardware
4. Monitor memory leaks

### Slashing Occurred

**Symptom**: Node got slashed and lost stake

**Solutions**:
1. Review slashing event logs
2. Fix the underlying issue (downtime/incorrect proofs)
3. Improve monitoring
4. Consider hardware/network upgrades

## Getting Help

- Check logs: `journalctl -u mindcelld -f`
- Discord: [community link]
- GitHub Issues: Report bugs
- Documentation: Read the full docs

## Debug Mode

Enable debug mode for detailed logging:

```bash
mindcelld start --log-level=debug
```

## Health Check

Check node health:

```bash
curl localhost:26657/health
```

