# Reward Module

## Overview

Distributes rewards to node operators based on their performance and contributions.

## Reward Sources

- Inference fees (30% of total)
- Block rewards (inflation)
- Performance bonuses

## Features

- Accumulated reward tracking
- Claim functionality
- Reward pool management
- Distribution logging

## Messages

- `MsgDistributeReward`: Allocate rewards to node
- `MsgClaimReward`: Claim pending rewards
- `MsgAddToPool`: Add funds to reward pool

## Claiming

Rewards can be claimed at any time. No minimum threshold.
