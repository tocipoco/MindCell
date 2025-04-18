#!/usr/bin/env bash

set -euo pipefail

echo "Initializing MindCell Testnet"
echo "=============================="

# Configuration
CHAIN_ID="${CHAIN_ID:-mindcell-testnet-1}"
VALIDATOR_NAME="${VALIDATOR_NAME:-validator}"
INITIAL_SUPPLY="1000000000mcell"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo "Chain ID: $CHAIN_ID"
echo "Validator: $VALIDATOR_NAME"

# Clean previous testnet data
echo ""
echo "Cleaning previous testnet data..."
rm -rf ~/.mindcell

# Initialize chain
echo ""
echo "Initializing chain..."
mindcelld init "$VALIDATOR_NAME" --chain-id "$CHAIN_ID"

# Create validator key
echo ""
echo "Creating validator key..."
mindcelld keys add "$VALIDATOR_NAME" --keyring-backend=test 2>&1 | tee validator_key.txt

VALIDATOR_ADDR=$(mindcelld keys show "$VALIDATOR_NAME" -a --keyring-backend=test)
echo "Validator address: $VALIDATOR_ADDR"

# Add genesis account
echo ""
echo "Adding genesis account..."
mindcelld add-genesis-account "$VALIDATOR_ADDR" "$INITIAL_SUPPLY"

# Create genesis transaction
echo ""
echo "Creating genesis transaction..."
mindcelld gentx "$VALIDATOR_NAME" \
    100000mcell \
    --chain-id="$CHAIN_ID" \
    --moniker="$VALIDATOR_NAME" \
    --commission-rate=0.10 \
    --commission-max-rate=0.20 \
    --commission-max-change-rate=0.01 \
    --min-self-delegation=1 \
    --keyring-backend=test

# Collect genesis transactions
echo ""
echo "Collecting genesis transactions..."
mindcelld collect-gentxs

# Validate genesis
echo ""
echo "Validating genesis..."
mindcelld validate-genesis

# Update configuration
echo ""
echo "Updating configuration..."

# Enable API
sed -i.bak 's/enable = false/enable = true/' "$HOME/.mindcell/config/app.toml"

# Set minimum gas prices
sed -i.bak 's/minimum-gas-prices = ""/minimum-gas-prices = "0.001mcell"/' "$HOME/.mindcell/config/app.toml"

# Enable unsafe CORS (testnet only)
sed -i.bak 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' "$HOME/.mindcell/config/app.toml"

echo ""
echo -e "${GREEN}✓ Testnet initialized successfully!${NC}"
echo ""
echo "Start the chain with:"
echo "  mindcelld start"
echo ""
echo "Validator address: $VALIDATOR_ADDR"
echo "Validator key saved to: validator_key.txt"
echo ""
echo -e "${YELLOW}⚠ Keep validator_key.txt secure!${NC}"
