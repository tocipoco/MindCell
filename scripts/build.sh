#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "Building MindCell..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check Go installation
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed${NC}"
    echo "Please install Go 1.21 or higher"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "Go version: $GO_VERSION"

# Clean previous builds
echo "Cleaning previous builds..."
rm -rf "$PROJECT_ROOT/build"
rm -rf "$PROJECT_ROOT/dist"

# Create build directory
mkdir -p "$PROJECT_ROOT/build"

# Get version from git
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')

echo "Building version: $VERSION"
echo "Commit: $COMMIT"
echo "Build time: $BUILD_TIME"

# Build flags
LDFLAGS="-X github.com/tocipoco/MindCell/app.Version=$VERSION"
LDFLAGS="$LDFLAGS -X github.com/tocipoco/MindCell/app.Commit=$COMMIT"
LDFLAGS="$LDFLAGS -X github.com/tocipoco/MindCell/app.BuildTime=$BUILD_TIME"

# Build for current platform
echo -e "${YELLOW}Building mindcelld binary...${NC}"
cd "$PROJECT_ROOT"

go build \
    -mod=readonly \
    -ldflags "$LDFLAGS" \
    -o build/mindcelld \
    ./cmd/mindcelld

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Build successful${NC}"
    echo "Binary location: $PROJECT_ROOT/build/mindcelld"
    
    # Display binary info
    ls -lh "$PROJECT_ROOT/build/mindcelld"
    "$PROJECT_ROOT/build/mindcelld" version
else
    echo -e "${RED}✗ Build failed${NC}"
    exit 1
fi

# Optional: Build for multiple platforms
if [ "${BUILD_ALL:-false}" = "true" ]; then
    echo -e "${YELLOW}Building for multiple platforms...${NC}"
    
    PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")
    
    for PLATFORM in "${PLATFORMS[@]}"; do
        GOOS=${PLATFORM%/*}
        GOARCH=${PLATFORM#*/}
        OUTPUT="build/mindcelld-$GOOS-$GOARCH"
        
        if [ "$GOOS" = "windows" ]; then
            OUTPUT+=".exe"
        fi
        
        echo "Building for $GOOS/$GOARCH..."
        GOOS=$GOOS GOARCH=$GOARCH go build \
            -mod=readonly \
            -ldflags "$LDFLAGS" \
            -o "$OUTPUT" \
            ./cmd/mindcelld
    done
    
    echo -e "${GREEN}✓ All platforms built${NC}"
    ls -lh "$PROJECT_ROOT/build/"
fi

echo -e "${GREEN}Build complete!${NC}"
