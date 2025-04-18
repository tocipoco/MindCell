#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "MindCell Release Script"
echo "======================="

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Parse version
VERSION="${1:-}"
if [ -z "$VERSION" ]; then
    echo -e "${RED}Error: Version required${NC}"
    echo "Usage: ./release.sh v1.0.0"
    exit 1
fi

# Validate version format
if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo -e "${RED}Error: Invalid version format${NC}"
    echo "Expected format: v1.0.0"
    exit 1
fi

echo "Preparing release: $VERSION"

cd "$PROJECT_ROOT"

# Check for uncommitted changes
if [ -n "$(git status --porcelain)" ]; then
    echo -e "${RED}Error: Uncommitted changes detected${NC}"
    echo "Please commit or stash changes before releasing"
    exit 1
fi

# Update to latest main
echo "Updating main branch..."
git checkout main
git pull origin main

# Run tests
echo ""
echo "Running test suite..."
./scripts/test.sh || {
    echo -e "${RED}✗ Tests failed${NC}"
    exit 1
}

# Run linters
echo ""
echo "Running linters..."
./scripts/lint.sh || {
    echo -e "${RED}✗ Linting failed${NC}"
    exit 1
}

# Update version in code
echo ""
echo "Updating version strings..."
sed -i.bak "s/const Version = .*/const Version = \"$VERSION\"/" app/version.go
rm -f app/version.go.bak

# Update CHANGELOG
echo ""
echo "Updating CHANGELOG..."
DATE=$(date '+%Y-%m-%d')
cat > CHANGELOG.tmp <<EOF
# Changelog

## $VERSION - $DATE

### Added
- New features and enhancements

### Changed
- Improvements and optimizations

### Fixed
- Bug fixes

---

$(cat CHANGELOG.md | tail -n +2)
EOF
mv CHANGELOG.tmp CHANGELOG.md

# Commit version bump
git add app/version.go CHANGELOG.md
git commit -m "chore: bump version to $VERSION"

# Create git tag
echo ""
echo "Creating git tag..."
git tag -a "$VERSION" -m "Release $VERSION"

# Build release binaries
echo ""
echo "Building release binaries..."
BUILD_ALL=true ./scripts/build.sh

# Create release archive
echo ""
echo "Creating release archives..."
mkdir -p dist

for BUILD in build/mindcelld-*; do
    if [ -f "$BUILD" ]; then
        PLATFORM=$(basename "$BUILD" | sed 's/mindcelld-//')
        ARCHIVE="dist/mindcell-$VERSION-$PLATFORM.tar.gz"
        
        echo "Creating $ARCHIVE..."
        tar -czf "$ARCHIVE" -C build "$(basename "$BUILD")" \
            -C "$PROJECT_ROOT" LICENSE README.md
        
        # Generate checksum
        sha256sum "$ARCHIVE" > "$ARCHIVE.sha256"
    fi
done

# Display artifacts
echo ""
echo "Release artifacts:"
ls -lh dist/

echo ""
echo -e "${GREEN}✓ Release $VERSION prepared successfully${NC}"
echo ""
echo "Next steps:"
echo "1. Review CHANGELOG.md"
echo "2. Push tag: git push origin $VERSION"
echo "3. Push commits: git push origin main"
echo "4. Create GitHub release with artifacts in dist/"
echo "5. Announce release to community"
