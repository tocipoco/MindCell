# MindCell Architecture

## Overview

MindCell is built on Cosmos SDK and Ethermint, providing a decentralized infrastructure for AI model sharding and inference.

## Core Components

### Model Registry
Manages AI model registration, versioning, and metadata storage.

### Shard Allocator
Distributes model shards across validator nodes using reputation-based allocation.

### Inference Gateway
Routes inference requests and verifies zkML proofs for correctness.

### Billing Module
Handles pay-per-inference fee calculation and distribution.

### Reward Module
Distributes rewards to node operators based on performance.

### Slashing Module
Penalizes misbehaving nodes through stake slashing.

### Token Module
Manages MCELL token supply and staking operations.

## Data Flow

1. Model owner registers a model
2. Shards are allocated to validator nodes
3. Client submits inference request
4. Nodes compute partial results
5. zkML proof is generated and verified
6. Payment is processed and distributed
7. Rewards are distributed to participants

