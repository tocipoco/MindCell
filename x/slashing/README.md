# Slashing Module

## Overview

Implements penalty mechanisms for misbehaving nodes to ensure network security and reliability.

## Slashing Conditions

| Violation | Penalty | Description |
|-----------|---------|-------------|
| Timeout | 5% | Node fails to respond in time |
| Incorrect Proof | 20% | zkML proof verification fails |
| Downtime | 10% | Extended unavailability |

## Parameters

- TimeoutSlashPercent: 0.05
- IncorrectProofSlashPercent: 0.20
- DowntimeSlashPercent: 0.10
- MinSlashAmount: 100 mcell
- MaxSlashAmount: 10,000 mcell

## Messages

- `MsgSlashNode`: Execute slashing penalty
- `MsgUpdateSlashingParams`: Update parameters (governance)

## Records

All slashing events are permanently recorded on-chain for transparency.
