# templates

A Go project that works with the [habits](https://github.com/sonaneir/habits) gRPC service, using its Protocol Buffers contract to communicate over gRPC. It builds on the generated gRPC client and adds tests with mocks, exercising the service's API from the client side.

## Overview

This module imports the `habits` gRPC service as a local dependency and interacts with it through its protobuf-defined contract. The focus is on the client side of gRPC — calling the service methods, handling the typed requests and responses generated from the `.proto` definitions, and testing that interaction with generated mocks.

## Tech stack

- **Go**
- **gRPC** (`google.golang.org/grpc`) — communication with the habits service
- **Protocol Buffers** (`google.golang.org/protobuf`) — the typed contract
- **minimock** (`github.com/gojuno/minimock`) — generated mocks for testing
- **testify** — test assertions

## Project structure

```
templates/
├── cmd/                    # entry point(s)
└── internal/               # client logic and tests
```

The `habits` service is wired in as a local module replacement, so the two projects are developed together — `templates` acts against the habits gRPC contract.

## What I learned

This project deepened my experience with gRPC and Protocol Buffers from the client's perspective — working with a typed contract, calling a gRPC service, and testing that communication with mock-based tests using minimock and testify. It's the kind of client-service interaction that underpins microservice communication.
