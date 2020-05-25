# Go lang web server + gRPC + Dgraph Database + React Typescript web client + Envoy

## Prerequisites

- docker
- [docker-compose](https://github.com/docker/compose)
## Instructions

1. Run _docker-compose_
```
docker-compose up --build
```

2. [Dgraph](https://github.com/dgraph-io/dgraph) can be accessed in _localhost:8000_. Go to this url and select _Latest_ or go directly to _localhost:8000/?latest_. Go to _Console_, then click _Mutate_.

3. Copy/Paste the contents of _mutate.json_ to the _Mutate_ box.

4. Query the data by copy/pasting the contents of _query.txt_ to the _Query_ box.

4. [React](https://github.com/facebook/react) is in _localhost:3000_

## Notes

- The _React_ portion is created using the [create-react-app](https://github.com/facebook/create-react-app) tool.

- [gRPC](https://github.com/grpc/grpc) server is listening on port _9090_.

- The _gen-proto_ file is a script to generate the proto file for _go_ and _typescript/js_.
