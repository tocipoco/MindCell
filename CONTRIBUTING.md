# Contributing to MindCell

Thank you for your interest in contributing to MindCell! This document provides guidelines for contributing to the project.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) before contributing.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- Make
- Basic understanding of Cosmos SDK
- Familiarity with blockchain concepts

### Setting Up Development Environment

```bash
# Fork and clone the repository
git clone https://github.com/YOUR_USERNAME/MindCell.git
cd MindCell

# Add upstream remote
git remote add upstream https://github.com/tocipoco/MindCell.git

# Install dependencies
make install

# Run tests
make test
```

## Development Workflow

### 1. Create a Branch

```bash
# Update main branch
git checkout main
git pull upstream main

# Create feature branch
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/bug-description
```

### 2. Make Changes

- Write clean, idiomatic Go code
- Follow existing code style and patterns
- Add tests for new functionality
- Update documentation as needed
- Keep commits logical and atomic

### 3. Test Your Changes

```bash
# Run all tests
make test

# Run specific module tests
go test ./x/modelregistry/...

# Check test coverage
make test-coverage

# Run linters
make lint

# Format code
make format
```

### 4. Commit Your Changes

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```bash
# Format
<type>: <description>

[optional body]

[optional footer]
```

**Types:**
- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `test:` Adding or updating tests
- `refactor:` Code refactoring
- `perf:` Performance improvements
- `chore:` Maintenance tasks
- `ci:` CI/CD changes

**Examples:**
```bash
git commit -m "feat: add model versioning support"
git commit -m "fix: correct shard allocation race condition"
git commit -m "docs: update API reference for billing module"
git commit -m "test: add integration tests for inference gateway"
```

### 5. Push and Create Pull Request

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
# Fill out the PR template with detailed information
```

## Pull Request Guidelines

### Before Submitting

- [ ] Code builds without errors
- [ ] All tests pass
- [ ] Code is properly formatted (`make format`)
- [ ] No linting errors (`make lint`)
- [ ] Documentation updated (if applicable)
- [ ] CHANGELOG.md updated (for significant changes)

### PR Description Should Include

1. **Summary**: What does this PR do?
2. **Motivation**: Why is this change necessary?
3. **Implementation**: How does it work?
4. **Testing**: How was it tested?
5. **Breaking Changes**: Any API changes?
6. **Related Issues**: Link to related issues

### Example PR Description

```markdown
## Summary
Implements automatic shard rebalancing based on node load.

## Motivation
Nodes were becoming overloaded while others sat idle. This feature
distributes load more evenly across the network.

## Implementation
- Added load monitoring to ShardAllocator keeper
- Implemented rebalancing algorithm based on node capacity
- Triggers rebalance every 1000 blocks if imbalance detected

## Testing
- Unit tests for rebalancing logic
- Integration tests simulating uneven load
- Testnet deployment verified

## Breaking Changes
None

## Related Issues
Closes #123
Related to #456
```

## Code Style Guide

### General Principles

1. **Clarity over cleverness**: Write code that's easy to understand
2. **Comments**: Explain why, not what
3. **Error handling**: Always handle errors appropriately
4. **Testing**: Write tests for all new code
5. **Documentation**: Update docs for user-facing changes

### Go Conventions

```go
// Good: Clear variable names
func (k Keeper) GetModel(ctx sdk.Context, modelID uint64) (types.Model, bool) {
    store := ctx.KVStore(k.storeKey)
    // ...
}

// Bad: Unclear abbreviations
func (k Keeper) GM(c sdk.Context, mid uint64) (types.Model, bool) {
    s := c.KVStore(k.sk)
    // ...
}

// Good: Proper error handling
func (k Keeper) ProcessInference(ctx sdk.Context, req InferenceRequest) error {
    if err := validateRequest(req); err != nil {
        return fmt.Errorf("invalid request: %w", err)
    }
    // ...
}

// Bad: Ignoring errors
func (k Keeper) ProcessInference(ctx sdk.Context, req InferenceRequest) {
    validateRequest(req) // error ignored
    // ...
}
```

### Testing Guidelines

```go
// Good: Table-driven tests
func TestModelRegistration(t *testing.T) {
    tests := []struct {
        name      string
        msg       types.MsgRegisterModel
        expectErr bool
    }{
        {
            name: "valid model",
            msg:  validMsg,
            expectErr: false,
        },
        {
            name: "invalid CID",
            msg:  invalidCIDMsg,
            expectErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic
        })
    }
}

// Good: Clear assertions
func TestGetModel(t *testing.T) {
    keeper, ctx := setupKeeper(t)
    
    model := types.Model{ID: 1, Owner: "cosmos1..."}
    keeper.SetModel(ctx, model)
    
    retrieved, found := keeper.GetModel(ctx, 1)
    require.True(t, found, "model should be found")
    require.Equal(t, model.ID, retrieved.ID, "model ID should match")
}
```

## Documentation

### Code Comments

```go
// GetModel retrieves a model by its ID.
// Returns the model and true if found, otherwise returns empty model and false.
func (k Keeper) GetModel(ctx sdk.Context, modelID uint64) (types.Model, bool) {
    // Implementation
}
```

### README Updates

- Keep README.md up to date with major features
- Add examples for new functionality
- Update installation instructions if dependencies change

### API Documentation

- Document all exported functions
- Include examples in godoc comments
- Update docs/api.md for user-facing changes

## Review Process

### What to Expect

1. **Initial Review** (1-3 days)
   - Automated checks run (CI/CD)
   - Maintainer does initial review

2. **Feedback Loop**
   - Address reviewer comments
   - Push updates to your branch
   - Discussion and iteration

3. **Approval**
   - At least one maintainer approval required
   - All checks must pass
   - No unresolved conversations

4. **Merge**
   - Squash and merge for clean history
   - Delete feature branch after merge

### Responding to Feedback

- Be respectful and professional
- Ask questions if feedback is unclear
- Explain your reasoning if you disagree
- Update code based on suggestions
- Mark conversations as resolved when addressed

## Common Tasks

### Adding a New Module

1. Create module directory structure
2. Implement types, keeper, and handlers
3. Add module to app.go
4. Write comprehensive tests
5. Document the module
6. Update go.mod if new dependencies

### Fixing a Bug

1. Write a failing test that reproduces the bug
2. Fix the bug
3. Verify the test now passes
4. Add regression test if applicable

### Updating Dependencies

```bash
# Update specific dependency
go get github.com/cosmos/cosmos-sdk@v0.50.2

# Update all dependencies
go get -u ./...

# Tidy and verify
go mod tidy
go mod verify
```

## Communication

### Where to Ask Questions

- **Discord**: Real-time chat and quick questions
- **GitHub Discussions**: Design discussions and proposals
- **GitHub Issues**: Bug reports and feature requests

### Reporting Bugs

Use the bug report template and include:
- Clear description of the issue
- Steps to reproduce
- Expected vs actual behavior
- Environment details
- Relevant logs or screenshots

### Proposing Features

Use the feature request template and include:
- Problem being solved
- Proposed solution
- Alternative approaches considered
- Impact on existing functionality

## Release Process

Maintainers handle releases:

1. Version bump in code
2. Update CHANGELOG.md
3. Create release tag
4. Build release binaries
5. Publish GitHub release
6. Announce to community

## Recognition

Contributors are recognized in:
- AUTHORS.md file
- Release notes
- Community shoutouts

Significant contributions may lead to:
- Commit access
- Maintainer status
- Core team invitation

## Questions?

If you have questions about contributing:
- Check existing issues and discussions
- Ask in Discord #development channel
- Email: dev@mindcell.network

Thank you for contributing to MindCell! 🚀
