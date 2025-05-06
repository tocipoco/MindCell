# Changelog

All notable changes to MindCell will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned Features
- FHE-based private inference
- Cross-chain inference via IBC
- Model NFT marketplace
- Dataset attribution system

## [1.0.0] - 2025-06-25

### Added
- ModelRegistry module for AI model registration and versioning
- ShardAllocator module with reputation-based node selection
- InferenceGateway module with zkML proof verification
- Billing module with multi-party revenue distribution
- Reward module for node operator incentives
- Slashing module for misbehavior penalties
- Token module for MCELL token management
- Complete CLI interface with all module commands
- Comprehensive test suite with >60% coverage
- Production-ready Docker images and compose files
- Extensive documentation covering all aspects
- CI/CD pipelines with automated testing
- Example configurations for various deployment scenarios

### Technical Details
- Built on Cosmos SDK v0.50.1
- Integrated Ethermint for EVM compatibility
- Support for gnark and Halo2 proof systems
- IPFS/Arweave integration for metadata storage
- Tendermint consensus with 3-second block time

### Documentation
- Complete API reference
- Deployment guides for validators and node operators
- Client SDK documentation (Go, JavaScript, Python)
- Performance optimization guides
- Security best practices
- Troubleshooting guides

## [0.1.0] - 2024-10-01

### Added
- Initial project structure
- Basic Cosmos SDK integration
- Module scaffolding
- Development environment setup

---

For a complete list of changes, see the [commit history](https://github.com/tocipoco/MindCell/commits/main/).
