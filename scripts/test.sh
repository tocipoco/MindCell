#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "MindCell Test Suite"
echo "==================="

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Configuration
COVERAGE_FILE="${COVERAGE_FILE:-coverage.txt}"
MIN_COVERAGE="${MIN_COVERAGE:-60}"
TEST_TIMEOUT="${TEST_TIMEOUT:-10m}"

cd "$PROJECT_ROOT"

# Check if go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed${NC}"
    exit 1
fi

# Clean previous coverage data
rm -f "$COVERAGE_FILE"

echo "Running tests with race detector and coverage..."
echo ""

# Run tests
if go test ./... \
    -v \
    -race \
    -timeout="$TEST_TIMEOUT" \
    -coverprofile="$COVERAGE_FILE" \
    -covermode=atomic \
    -count=1; then
    
    echo ""
    echo -e "${GREEN}✓ All tests passed${NC}"
    
    # Display coverage summary
    echo ""
    echo "Coverage Summary:"
    echo "================"
    go tool cover -func="$COVERAGE_FILE" | tail -10
    
    # Check coverage threshold
    TOTAL_COVERAGE=$(go tool cover -func="$COVERAGE_FILE" | grep total | awk '{print $3}' | sed 's/%//')
    
    if (( $(echo "$TOTAL_COVERAGE >= $MIN_COVERAGE" | bc -l) )); then
        echo -e "${GREEN}✓ Coverage ($TOTAL_COVERAGE%) meets minimum threshold ($MIN_COVERAGE%)${NC}"
    else
        echo -e "${YELLOW}⚠ Coverage ($TOTAL_COVERAGE%) below minimum threshold ($MIN_COVERAGE%)${NC}"
        exit 1
    fi
    
    # Generate HTML coverage report
    if [ "${HTML_COVERAGE:-false}" = "true" ]; then
        echo ""
        echo "Generating HTML coverage report..."
        go tool cover -html="$COVERAGE_FILE" -o coverage.html
        echo "Report generated: coverage.html"
    fi
    
else
    echo ""
    echo -e "${RED}✗ Tests failed${NC}"
    exit 1
fi

# Run specific module tests if requested
if [ -n "${TEST_MODULE:-}" ]; then
    echo ""
    echo "Running tests for module: $TEST_MODULE"
    go test "./x/$TEST_MODULE/..." -v -race
fi

# Run benchmark tests if requested
if [ "${BENCH:-false}" = "true" ]; then
    echo ""
    echo "Running benchmark tests..."
    go test -bench=. -benchmem ./... | tee benchmark.txt
fi

# Run integration tests if requested
if [ "${INTEGRATION:-false}" = "true" ]; then
    echo ""
    echo "Running integration tests..."
    go test -tags=integration ./... -v
fi

echo ""
echo -e "${GREEN}Test suite completed successfully!${NC}"

