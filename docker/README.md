# Docker Deployment Guide

## Quick Start

### Using Docker Compose (Recommended)

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f mindcell

# Stop services
docker-compose down
```

This starts:
- MindCell node
- Prometheus (metrics)
- Grafana (dashboards)

### Using Docker Only

```bash
# Build image
docker build -t mindcell:latest .

# Run container
docker run -d \
  --name mindcell-node \
  -p 26656:26656 \
  -p 26657:26657 \
  -p 1317:1317 \
  -v mindcell-data:/home/mindcell/.mindcell \
  mindcell:latest
```

## Configuration

### Environment Variables

```bash
CHAIN_ID=mindcell-1          # Chain identifier
MONIKER=my-node              # Node name
SEEDS=seed1:26656,seed2:26656  # Seed nodes
PERSISTENT_PEERS=peer1:26656   # Persistent peers
MINIMUM_GAS_PRICES=0.001mcell  # Min gas price
```

### Volume Mounts

- `/home/mindcell/.mindcell` - Node data directory
- `/home/mindcell/config` - Configuration files (optional)

## Monitoring Stack

### Access Services

- **Node RPC**: http://localhost:26657
- **REST API**: http://localhost:1317
- **Prometheus**: http://localhost:9091
- **Grafana**: http://localhost:3000 (admin/admin)

### Grafana Dashboards

Pre-configured dashboards:
- Node health and performance
- Transaction metrics
- Network statistics
- Resource utilization

## Production Deployment

### Using Docker Swarm

```bash
# Initialize swarm
docker swarm init

# Deploy stack
docker stack deploy -c docker-compose.yml mindcell

# Scale services
docker service scale mindcell_mindcell=3
```

### Using Kubernetes

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mindcell
spec:
  serviceName: mindcell
  replicas: 3
  selector:
    matchLabels:
      app: mindcell
  template:
    metadata:
      labels:
        app: mindcell
    spec:
      containers:
      - name: mindcell
        image: mindcell:latest
        ports:
        - containerPort: 26656
        - containerPort: 26657
        volumeMounts:
        - name: data
          mountPath: /home/mindcell/.mindcell
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 500Gi
```

## Backup and Recovery

### Backup Node Data

```bash
# Create backup
docker run --rm \
  -v mindcell-data:/data \
  -v $(pwd):/backup \
  alpine tar czf /backup/mindcell-backup.tar.gz /data

# Restore backup
docker run --rm \
  -v mindcell-data:/data \
  -v $(pwd):/backup \
  alpine sh -c "cd /data && tar xzf /backup/mindcell-backup.tar.gz --strip 1"
```

## Troubleshooting

### Container Won't Start

```bash
# Check logs
docker logs mindcell-node

# Inspect container
docker inspect mindcell-node

# Enter container
docker exec -it mindcell-node sh
```

### Reset Node State

```bash
# Stop and remove container
docker-compose down

# Remove volumes
docker volume rm mindcell_mindcell-data

# Restart
docker-compose up -d
```

## Security

- Run as non-root user (mindcell:1000)
- Read-only root filesystem where possible
- Limited capabilities
- Network isolation
- Secret management via Docker secrets or Kubernetes secrets