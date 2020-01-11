#!/usr/bin/env bash
set -euxo pipefail
mkdir -p $(pwd)/docs/proto
docker run --rm -v $(pwd)/api:/protos \
 -v $(pwd)/docs/proto:/out \
 pseudomuto/protoc-gen-doc
