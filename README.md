# Becca's Go Boilerplate
This implements a service like that described in the gRPC examples,
where a client provides their name and the service responds with a 
personalized greeting.

The primary difference is that this contains a lot of cleanup that I
find makes it easier to dive into.

## TL;DR
- Go gRPC Hello World service
- Docker-based protobuf compilation pipeline
- Protobuf documentation generator

## Paths
`scripts/` - General scripted tasks

`cmd/` - All services that will yield binaries

`api/` - Protobuf files

`internal/` - Internal code, incl. generated protobufs when built

`docs/` - User-written docs, incl. generated protobuf docs when built

### Generated Paths
`internal/proto/` - Generated protobuf files, not committed to git

`docs/proto/` - Generated protobuf docs, not committed to git

## Common Tasks
### Compiling Protocol Buffers
Use the  `scripts/generate-proto.sh` script from this directory.
```
./scripts/generate-proto.sh
```
### Generating Protobuf Docs
Use the  `scripts/generate-proto-docs.sh` script from this directory.
```
./scripts/generate-proto-docs.sh
```