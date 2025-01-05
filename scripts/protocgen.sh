#!/usr/bin/env bash

set -eo pipefail

echo "Generating protobuf files..."

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    echo "Processing $file"
  done
done

echo "Proto generation complete"

