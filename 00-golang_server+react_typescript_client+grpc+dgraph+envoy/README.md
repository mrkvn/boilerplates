# Go lang web server + gRPC + Dgraph Database + React Typescript web client + Envoy

## Prerequisites

- docker
- [docker-compose](https://github.com/docker/compose)
## Instructions

1. Run _docker-compose_
```
docker-compose up --build
```

2. [Dgraph](https://github.com/dgraph-io/dgraph) can be accessed in _localhost:8080_

3. [React](https://github.com/facebook/react) is in _localhost:3000_

## Notes

- The _React_ portion is created using the [create-react-app](https://github.com/facebook/create-react-app) tool.

- [gRPC](https://github.com/grpc/grpc) server is listening on port _9090_.

- The _gen-proto_ file is a script to generate the proto file for _go_ and _typescript/js_.
