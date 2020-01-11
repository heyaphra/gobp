#!/usr/bin/env bash
set -euxo pipefail
mkdir -p $(pwd)/internal/proto
docker run --rm -v $(pwd)/api:/protobuf \
 -v $(pwd)/internal/proto:/output \
 namely/protoc-all \
 -o /output \
 -d /protobuf \
 -l go
