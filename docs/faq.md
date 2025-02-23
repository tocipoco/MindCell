# Frequently Asked Questions

## General

### What is MindCell?
MindCell is a decentralized protocol for AI model sharding, inference, and monetization.

### What blockchain does MindCell use?
MindCell is built on Cosmos SDK with Ethermint compatibility.

### What is the native token?
MCELL is the native token used for staking, gas, and inference payments.

## For Model Owners

### How do I register my model?
Use the ModelRegistry module to register your model with metadata stored on IPFS.

### How do I earn from my model?
You earn 60% of inference fees when users query your model.

### Can I update my model?
Yes, you can update model metadata and create new versions.

## For Node Operators

### What are the hardware requirements?
Minimum: 8 CPU cores, 16GB RAM, 500GB SSD, stable network.

### How much stake is required?
Minimum stake is 1,000 MCELL to become a node operator.

### How do I claim rewards?
Use the reward claiming transaction to withdraw accumulated earnings.

### What causes slashing?
Slashing occurs for timeouts, incorrect proofs, or extended downtime.

## For Users

### How do I submit an inference request?
Use the InferenceGateway to submit requests with your input data and payment.

### How are fees calculated?
Fees = baseFee + (computeUnits × computePrice)

### Are my inference requests private?
Basic requests are not private, but FHE-enabled privacy mode is planned.

## Technical

### What is zkML?
Zero-knowledge machine learning ensures inference correctness without revealing the model.

### How does shard allocation work?
Shards are allocated based on node capacity, reputation, and network topology.

### Is the code open source?
Yes, MindCell is fully open source under MIT license.

