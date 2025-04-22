FROM golang:1.21-alpine AS builder

# Install dependencies
RUN apk add --no-cache git make gcc musl-dev linux-headers

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN make build

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates bash

# Copy binary from builder
COPY --from=builder /app/build/mindcelld /usr/local/bin/

# Create mindcell user
RUN addgroup -g 1000 mindcell && \
    adduser -D -u 1000 -G mindcell mindcell

# Set working directory
WORKDIR /home/mindcell

# Switch to mindcell user
USER mindcell

# Expose ports
EXPOSE 26656 26657 26660 1317 9090

# Health check
HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD mindcelld status 2>&1 | jq -r '.SyncInfo.catching_up' | grep -q false

# Default command
CMD ["mindcelld", "start"]
