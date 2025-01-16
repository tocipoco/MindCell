# Getting Started with MindCell

## Prerequisites

- Go 1.21 or higher
- Cosmos SDK v0.50+
- Make

## Installation

```bash
# Clone the repository
git clone https://github.com/tocipoco/MindCell.git
cd MindCell

# Install dependencies
./scripts/install-deps.sh

# Build the binary
make install
```

## Running a Node

```bash
# Initialize the node
mindcelld init mynode --chain-id mindcell-1

# Start the node
mindcelld start
```

## Configuration

Edit the configuration file at `~/.mindcell/config/config.toml` to customize your node settings.

## Next Steps

- Read the [Architecture](architecture.md) document
- Explore the [API Documentation](api.md)
- Join our community

