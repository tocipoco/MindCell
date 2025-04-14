# Governance Guide

## Overview

MindCell uses on-chain governance for protocol upgrades, parameter changes, and treasury management.

## Governance Token

**MCELL** token holders can:
- Submit proposals
- Vote on active proposals
- Delegate voting power to validators
- Earn governance rewards

## Proposal Types

### 1. Text Proposals

General signaling or coordination proposals without code changes.

**Example:**
```bash
mindcelld tx gov submit-proposal \
  --title="Increase Marketing Budget" \
  --description="Proposal to allocate 100,000 MCELL for Q2 marketing" \
  --type="Text" \
  --deposit=1000mcell \
  --from=mykey
```

### 2. Parameter Change Proposals

Modify chain parameters without software upgrade.

**Common Parameters:**
- `billing/BaseFee`: Minimum fee per inference
- `billing/ModelOwnerPercent`: Revenue share for model owners
- `slashing/TimeoutSlashPercent`: Penalty for timeouts
- `token/MinStakeAmount`: Minimum stake requirement

**Example:**
```json
{
  "title": "Reduce Base Inference Fee",
  "description": "Lower base fee from 1000 to 500 mcell to increase adoption",
  "changes": [
    {
      "subspace": "billing",
      "key": "BaseFee",
      "value": "500"
    }
  ],
  "deposit": "10000mcell"
}
```

Submit:
```bash
mindcelld tx gov submit-proposal param-change proposal.json \
  --from=mykey \
  --chain-id=mindcell-1
```

### 3. Software Upgrade Proposals

Coordinate network-wide software upgrades.

**Example:**
```json
{
  "title": "Upgrade to v2.0.0",
  "description": "Major upgrade with FHE support and performance improvements",
  "plan": {
    "name": "v2.0.0",
    "height": 1000000,
    "info": "https://github.com/tocipoco/MindCell/releases/tag/v2.0.0"
  },
  "deposit": "50000mcell"
}
```

### 4. Community Pool Spend Proposals

Allocate funds from community treasury.

**Example:**
```bash
mindcelld tx gov submit-proposal community-pool-spend \
  --title="Grant for zkML Research" \
  --description="Fund research into advanced proof systems" \
  --recipient=cosmos1recipient... \
  --amount=100000mcell \
  --deposit=10000mcell \
  --from=mykey
```

## Proposal Lifecycle

### 1. Deposit Period (7 days)

- Proposal needs minimum deposit: **10,000 MCELL**
- Anyone can contribute to deposit
- If minimum not reached, proposal rejected
- Deposits refunded if proposal passes or is rejected
- Deposits burned if proposal is vetoed

### 2. Voting Period (14 days)

Once minimum deposit reached, voting begins:

**Voting Options:**
- **Yes**: Support the proposal
- **No**: Oppose the proposal  
- **Abstain**: Acknowledge without taking position
- **NoWithVeto**: Oppose and want deposit burned (for spam/malicious proposals)

**Quorum:** 40% of voting power must participate

**Pass Threshold:** 50% of non-abstain votes must be Yes

**Veto Threshold:** 33.4% NoWithVeto fails proposal and burns deposit

### 3. Execution

If proposal passes:
- Parameter changes: Applied immediately
- Software upgrades: Executed at specified block height
- Treasury spends: Funds transferred
- Text proposals: No on-chain action

## Voting Guide

### As Token Holder

```bash
# List active proposals
mindcelld query gov proposals --status=voting_period

# View proposal details
mindcelld query gov proposal 1

# Vote on proposal
mindcelld tx gov vote 1 yes \
  --from=mykey \
  --chain-id=mindcell-1 \
  --fees=100mcell

# Check your vote
mindcelld query gov vote 1 $(mindcelld keys show mykey -a)
```

### As Validator

Validators should:
1. Review proposals thoroughly
2. Discuss with community
3. Vote responsibly (delegators inherit vote)
4. Explain voting rationale publicly

**Validator vote:**
```bash
mindcelld tx gov vote 1 yes \
  --from=validator \
  --chain-id=mindcell-1
```

### Delegation

Delegate voting power to a validator:

```bash
# Your delegated MCELL inherits validator's vote
# Unless you vote directly (overrides delegation)

# Vote directly (overrides validator vote for your stake)
mindcelld tx gov vote 1 no \
  --from=delegator \
  --chain-id=mindcell-1
```

## Submitting Proposals

### Best Practices

1. **Community Discussion First**
   - Post in forums/Discord
   - Gather feedback
   - Revise based on input

2. **Clear Communication**
   - Concise title (< 100 chars)
   - Detailed description
   - Include rationale and impact analysis

3. **Technical Specifications**
   - For code changes, link to PR
   - Include testing results
   - Document risks and mitigations

4. **Timing**
   - Avoid major holidays
   - Give validators time to review
   - Coordinate with core team for upgrades

### Proposal Template

```markdown
# Title: [Concise description]

## Summary
Brief overview of the proposal (2-3 sentences)

## Motivation
Why is this change needed?

## Specification
Technical details of the change

## Impact Analysis
- Who benefits?
- Who might be negatively affected?
- Resource requirements?

## Implementation
- Timeline
- Responsible parties
- Success criteria

## Risks and Mitigations
Potential issues and how to address them

## Alternatives Considered
Other approaches and why they were rejected

## References
- Related proposals
- Supporting documentation
- Community discussions
```

## Proposal Examples

### Example 1: Fee Reduction

```json
{
  "title": "Reduce Minimum Inference Fee by 50%",
  "description": "Current base fee of 1000mcell is deterring small-scale users. Reducing to 500mcell will increase accessibility while maintaining network sustainability. Analysis shows 3x increase in usage would compensate for lower fees.",
  "changes": [
    {
      "subspace": "billing",
      "key": "BaseFee",
      "value": "500"
    }
  ],
  "deposit": "10000mcell"
}
```

### Example 2: Slashing Parameter Update

```json
{
  "title": "Decrease Timeout Slash Percentage",
  "description": "Current 5% slash for timeouts is harsh given network latency variations. Propose reducing to 2% to be more lenient while maintaining accountability.",
  "changes": [
    {
      "subspace": "slashing",
      "key": "TimeoutSlashPercent",
      "value": "0.02"
    }
  ],
  "deposit": "10000mcell"
}
```

## Emergency Procedures

### Circuit Breaker

In case of critical bugs or attacks:

```bash
# Validators can halt the chain
mindcelld tx crisis invariant-broken \
  --from=validator \
  --invariant-route=module/route

# Requires coordination via validator chat
```

### Emergency Upgrade

For critical security fixes:

1. Core team prepares patch
2. Emergency proposal with 48-hour voting period
3. Coordinated upgrade at specific block height
4. Fallback plan if upgrade fails

## Governance Analytics

### Track Voting Activity

```bash
# See vote distribution
mindcelld query gov votes 1

# Tally results
mindcelld query gov tally 1
```

### Participation Metrics

Monitor on block explorer:
- Voter turnout by proposal
- Top voting addresses
- Validator voting records
- Proposal success rate

## Incentives

### Governance Rewards

- Voters earn 2% APY (from inflation)
- Rewards proportional to stake
- Only active voters qualify
- Claimed with regular staking rewards

### Reputation

- Consistent participation builds reputation
- May influence future proposals
- Validator selection by delegators

## Common Governance Questions

### How much deposit is required?

Minimum: 10,000 MCELL. Can be crowdfunded.

### Can I change my vote?

Yes, until voting period ends. Latest vote counts.

### What if I don't vote?

Your stake inherits your validator's vote (if delegated).

### When are deposits returned?

- If proposal passes: Immediately after execution
- If proposal rejected: Immediately
- If proposal vetoed (>33% NoWithVeto): Burned

### Can proposals be cancelled?

No, once submitted they must complete the process.

## Resources

- Governance forum: https://forum.mindcell.network
- Proposal tracker: https://explorer.mindcell.network/governance
- Voting guide: https://docs.mindcell.network/governance
- Discord: #governance channel
