# Migration Guide

## Upgrading from v0.x to v1.0

### Breaking Changes

1. **API Changes**
   - `RegisterModel` now requires `ShardCount` parameter
   - `SubmitInference` renamed from `RequestInference`
   - New required field: `MetadataCID` for models

2. **State Structure**
   - Model storage format changed
   - Requires state migration

### Migration Steps

#### 1. Export Current State

```bash
# Before upgrade, export state
mindcelld export > state_export.json
```

#### 2. Install New Version

```bash
# Download v1.0.0
wget https://github.com/tocipoco/MindCell/releases/download/v1.0.0/mindcelld
chmod +x mindcelld
sudo mv mindcelld /usr/local/bin/
```

#### 3. Migrate State

```bash
# Run migration script
mindcelld migrate v1.0.0 state_export.json --genesis-time=2025-06-01T00:00:00Z > new_genesis.json

# Replace genesis
cp new_genesis.json ~/.mindcell/config/genesis.json
```

#### 4. Reset Data

```bash
# Unsafe reset (only if coordinated)
mindcelld unsafe-reset-all

# Safe reset (recommended)
mindcelld tendermint unsafe-reset-all --keep-addr-book
```

#### 5. Restart Node

```bash
sudo systemctl start mindcelld
```

### Code Migration

**Old Code (v0.x):**
```javascript
await client.requestInference({
  modelHash: 'abc123',
  input: data
});
```

**New Code (v1.0):**
```javascript
await client.inference.submit({
  modelId: 42,
  inputData: data,
  maxFee: '500mcell'
});
```

## Database Migrations

### Add New Indices

If upgrade adds new query patterns:

```bash
# No manual intervention needed
# Cosmos SDK rebuilds indices automatically
# May take 10-30 minutes depending on state size
```

## Configuration Updates

### app.toml Changes

New configuration options in v1.0:

```toml
[mindcell]
enable-inference-cache = true
inference-cache-size = 1024
max-concurrent-inferences = 100
```

Add these to your `app.toml` after upgrade.

## Data Migration Scripts

For complex migrations, use provided scripts:

```bash
# Migrate model registry data
./scripts/migrate-models.sh state_export.json

# Verify migration
./scripts/verify-migration.sh
```

## Compatibility Matrix

| Client SDK | Node Version | Compatible |
|-----------|--------------|------------|
| 0.x       | 1.0          | ✗          |
| 1.x       | 1.0          | ✓          |
| 1.x       | 0.x          | ✗          |

## Rollback Plan

If community decides to rollback:

1. Governance proposal to revert
2. Coordinated halt
3. Restore pre-upgrade state
4. Downgrade binaries
5. Restart network

## Support

For migration help:
- Discord: #upgrades channel
- Email: upgrades@mindcell.network
