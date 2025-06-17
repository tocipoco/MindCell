# ShardAllocator Module

## Overview

Manages the distribution of model shards across validator nodes with reputation-based allocation.

## Key Features

- Node registration and management
- Intelligent shard assignment algorithm
- Reputation tracking
- Load balancing
- Shard replacement and reallocation

## State

- ShardAssignments: Shard to node mappings
- NodeReputations: Performance tracking
- RegisteredNodes: Active node registry

## Messages

- `MsgRegisterNode`: Register as shard hosting node
- `MsgAssignShard`: Assign shard to node
- `MsgReplaceShard`: Replace failed node assignment
- `MsgUpdateNodeReputation`: Update reputation score

## Node Selection Algorithm

Nodes selected based on:
- Available capacity
- Reputation score
- Geographic distribution
- Historical performance
