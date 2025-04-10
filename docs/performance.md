# Performance Optimization Guide

## Node Performance

### Hardware Recommendations

**Minimum Requirements:**
- CPU: 8 cores @ 2.5GHz
- RAM: 16GB DDR4
- Storage: 500GB NVMe SSD
- Network: 100Mbps symmetric

**Recommended for Production:**
- CPU: 16 cores @ 3.0GHz+
- RAM: 32GB DDR4
- Storage: 1TB NVMe SSD (RAID 1)
- Network: 1Gbps symmetric

### Operating System Tuning

```bash
# Increase file descriptor limits
ulimit -n 65536

# Optimize network buffers
sysctl -w net.core.rmem_max=16777216
sysctl -w net.core.wmem_max=16777216

# Disable swap for consistent performance
swapoff -a
```

### Database Optimization

**LevelDB Settings:**
```toml
[leveldb]
block_cache_size = 1073741824  # 1GB
write_buffer_size = 67108864    # 64MB
max_open_files = 1000
```

**Pruning Strategy:**
- Keep last 100,000 blocks
- Prune every 10,000 blocks
- Reduces disk usage by 70-80%

## Inference Optimization

### Batch Processing

Process multiple inference requests together:

```go
// Instead of processing one by one
for _, request := range requests {
    process(request)  // Slow
}

// Batch process for 10x speedup
batchSize := 32
processBatch(requests[:batchSize])
```

**Benefits:**
- 10x throughput improvement
- Better GPU utilization
- Reduced overhead per request

### Caching Strategies

**Model Metadata Cache:**
- Cache frequently accessed models
- TTL: 5 minutes
- Hit rate target: >95%

**Shard Location Cache:**
- Cache shard-to-node mappings
- Invalidate on reallocation
- Reduces query latency by 80%

### Connection Pooling

Maintain persistent connections to shard nodes:

```go
pool := NewConnectionPool(
    maxConnections: 100,
    idleTimeout: 30 * time.Second,
    maxLifetime: 10 * time.Minute,
)
```

## Network Optimization

### Peer Selection

- Prioritize geographically close peers
- Maintain 10-20 persistent connections
- Rotate seed nodes hourly

### Bandwidth Management

- Rate limit incoming connections: 50 conn/sec
- Prioritize consensus messages
- Compress large payloads (>1KB)

### Latency Reduction

**Target Metrics:**
- Block time: 3 seconds
- Transaction finality: 2 blocks (6 seconds)
- Inference response: <2 seconds (p95)

**Techniques:**
- Asynchronous proof generation
- Parallel shard queries
- Edge node deployment

## Memory Management

### Go Runtime Tuning

```bash
export GOGC=100          # Default GC target
export GOMEMLIMIT=28GiB  # Soft memory limit
```

### Memory Profiling

```bash
# Enable memory profiling
go tool pprof http://localhost:6060/debug/pprof/heap

# Identify leaks
go tool pprof -alloc_space mindcelld memory.prof
```

### Optimization Checklist

- [ ] Reuse buffers and objects
- [ ] Avoid unnecessary allocations in hot paths
- [ ] Use sync.Pool for frequently allocated objects
- [ ] Profile before optimizing

## Monitoring Dashboards

### Key Metrics to Track

**System Level:**
- CPU usage: <70% average
- Memory usage: <80% total
- Disk I/O: <50% capacity
- Network throughput

**Application Level:**
- Transactions per second (TPS)
- Block production time
- Consensus participation rate
- Peer count and connectivity

**Business Level:**
- Inference requests per hour
- Average inference latency
- Revenue per node
- Slashing events count

### Alerting Rules

```yaml
alerts:
  - name: HighLatency
    condition: p95_latency > 3s
    severity: warning
    
  - name: NodeDowntime
    condition: uptime < 99%
    severity: critical
    
  - name: HighMemoryUsage
    condition: memory_usage > 90%
    severity: warning
```

## Scalability Considerations

### Horizontal Scaling

- Add more shard nodes as demand grows
- Each node handles 50-100 shards optimally
- Network can support 10,000+ nodes theoretically

### Vertical Scaling

- Upgrade node hardware for better performance
- Diminishing returns above 32 cores
- Focus on SSD speed and network bandwidth

### Future Improvements

- Layer 2 solutions for high-frequency inference
- State channels for micro-payments
- Rollups for computation aggregation
- Sharding the blockchain itself (long-term)

## Benchmarking

### Standard Test Suite

```bash
# Run benchmark tests
go test -bench=. -benchmem ./...

# Specific module benchmarks
go test -bench=BenchmarkInference ./x/inferencegateway/...
```

### Expected Performance

| Operation | Throughput | Latency (p95) |
|-----------|-----------|---------------|
| Model registration | 100/min | 500ms |
| Shard assignment | 1000/min | 100ms |
| Inference submission | 10,000/min | 200ms |
| Proof verification | 5,000/min | 50ms |
| Reward distribution | 500/min | 300ms |

## Troubleshooting Performance Issues

### High CPU Usage

**Diagnosis:**
```bash
top -H -p $(pgrep mindcelld)
perf record -g -p $(pgrep mindcelld)
```

**Common Causes:**
- Inefficient proof verification
- Too many concurrent goroutines
- Busy consensus participation

**Solutions:**
- Increase worker pool size
- Implement backpressure mechanisms
- Optimize hot code paths

### High Memory Usage

**Diagnosis:**
```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap
```

**Common Causes:**
- Memory leaks in event handlers
- Unbounded cache growth
- Large transaction pools

**Solutions:**
- Implement LRU caches with size limits
- Regular garbage collection
- Monitor goroutine counts

### Slow Queries

**Diagnosis:**
- Enable query logging
- Use database profiling tools
- Check index utilization

**Solutions:**
- Add database indices on frequently queried fields
- Denormalize data for read-heavy workloads
- Implement query result caching

## Production Best Practices

1. **Monitor continuously**: Set up Prometheus + Grafana
2. **Test under load**: Use realistic traffic patterns
3. **Plan capacity**: Forecast growth and scale proactively
4. **Optimize iteratively**: Profile → identify bottlenecks → fix → repeat
5. **Document changes**: Track performance over time
6. **Review regularly**: Monthly performance audits

