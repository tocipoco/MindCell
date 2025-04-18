#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "MindCell Deployment Script"
echo "=========================="

# Configuration
DEPLOY_ENV="${DEPLOY_ENV:-production}"
NODE_NAME="${NODE_NAME:-mindcell-node}"
CHAIN_ID="${CHAIN_ID:-mindcell-1}"
MONIKER="${MONIKER:-$NODE_NAME}"

echo "Environment: $DEPLOY_ENV"
echo "Chain ID: $CHAIN_ID"
echo "Moniker: $MONIKER"

# Check if running as root
if [ "$EUID" -eq 0 ]; then
    echo "Error: Do not run this script as root"
    exit 1
fi

# Install binary
echo "Installing mindcelld binary..."
cd "$PROJECT_ROOT"
make install

# Verify installation
if ! command -v mindcelld &> /dev/null; then
    echo "Error: mindcelld not found in PATH"
    exit 1
fi

echo "✓ Binary installed: $(which mindcelld)"
mindcelld version

# Initialize node
if [ ! -d "$HOME/.mindcell" ]; then
    echo "Initializing node..."
    mindcelld init "$MONIKER" --chain-id "$CHAIN_ID"
else
    echo "Node already initialized at $HOME/.mindcell"
fi

# Download genesis file
if [ "$DEPLOY_ENV" = "production" ]; then
    echo "Downloading mainnet genesis..."
    curl -L https://raw.githubusercontent.com/mindcell-network/networks/main/mainnet/genesis.json \
        > "$HOME/.mindcell/config/genesis.json"
elif [ "$DEPLOY_ENV" = "testnet" ]; then
    echo "Downloading testnet genesis..."
    curl -L https://raw.githubusercontent.com/mindcell-network/networks/main/testnet/genesis.json \
        > "$HOME/.mindcell/config/genesis.json"
fi

# Configure seeds and peers
if [ "$DEPLOY_ENV" = "production" ]; then
    SEEDS="seed1.mindcell.network:26656,seed2.mindcell.network:26656"
    PEERS="peer1.mindcell.network:26656,peer2.mindcell.network:26656"
else
    SEEDS="testnet-seed.mindcell.network:26656"
    PEERS="testnet-peer.mindcell.network:26656"
fi

# Update config
sed -i.bak "s/^seeds =.*/seeds = \"$SEEDS\"/" "$HOME/.mindcell/config/config.toml"
sed -i.bak "s/^persistent_peers =.*/persistent_peers = \"$PEERS\"/" "$HOME/.mindcell/config/config.toml"

# Set minimum gas prices
sed -i.bak 's/^minimum-gas-prices =.*/minimum-gas-prices = "0.001mcell"/' "$HOME/.mindcell/config/app.toml"

# Enable API and gRPC
sed -i.bak 's/^enable = false/enable = true/' "$HOME/.mindcell/config/app.toml"

# Create systemd service
echo "Creating systemd service..."
sudo tee /etc/systemd/system/mindcelld.service > /dev/null <<EOF
[Unit]
Description=MindCell Node
After=network-online.target

[Service]
User=$USER
ExecStart=$(which mindcelld) start
Restart=always
RestartSec=3
LimitNOFILE=65535
Environment="DAEMON_HOME=$HOME/.mindcell"

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd
sudo systemctl daemon-reload
sudo systemctl enable mindcelld

echo ""
echo "Deployment complete!"
echo ""
echo "To start the node:"
echo "  sudo systemctl start mindcelld"
echo ""
echo "To view logs:"
echo "  journalctl -u mindcelld -f"
echo ""
echo "To check status:"
echo "  mindcelld status"
