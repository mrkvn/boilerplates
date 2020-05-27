# Go lang + gRPC + Dgraph Database + React Typescript + Envoy

## Prerequisites

- docker
- [docker-compose](https://github.com/docker/compose)
- [go](https://github.com/golang/go)
- [protoc](https://github.com/protocolbuffers/protobuf/)
- [protoc-gen-go](https://github.com/golang/protobuf/tree/master/protoc-gen-go)
	```
	go get github.com/golang/protobuf/protoc-gen-go
	```
- [grpc-web](https://github.com/grpc/grpc-web)
- [dgo](https://github.com/dgraph-io/dgo)
	```
	go get github.com/dgraph-io/dgo
	```
- [curl](https://github.com/curl/curl)
- [Chrome Extension: GraphiQL](https://chrome.google.com/webstore/detail/graphiql-extension/jhbedfdjpmemmbghfecnaeeiokonjclb?hl=en)

## Instructions

- Run _docker-compose_.
	```
	docker-compose up --build
	```


- To access [Dgraph](https://github.com/dgraph-io/dgraph)'s _GraphQL_ endpoint, upload a [schema](dgraph/schema.graphql) using [curl](https://github.com/curl/curl).
	```
	curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'
	```

- Open the [Chrome Extension: GraphiQL](https://chrome.google.com/webstore/detail/graphiql-extension/jhbedfdjpmemmbghfecnaeeiokonjclb?hl=en) and set the endpoint to _http://localhost:8080/graphql_. Copy/Paste the contents of [mutate](dgraph/mutate.txt) to the box provided then click the _Play_ button or press _Ctrl-Enter_.

	_Note: This is GraphQL syntax_

	```
	mutation {
		addProduct(input: [
			{ name: "GraphQL on Dgraph"},
			{ name: "Dgraph: The GraphQL Database"}
		]) {
			product {
				productID
				name
			}
		}
		addCustomer(input: [{ username: "Michael"}]) {
			customer {
				username
			}
		}
	}
	```

- Access Dgraph Ratel on _localhost:8000_ and select _Latest_.

- Query the data by copy/pasting the contents of [query](dgraph/query.txt) to the _Query_ box. Data can also be queried from GraphiQL though the syntax would be different (GraphQL syntax must be used).

	_Note: This is GraphQL+- syntax_

	```
	{
		customer(func: has(Customer.username)) {
			Customer.username
		}
	}
	```

- Open the [React](https://github.com/facebook/react) app (_localhost:3000_) and see the console log. Two queries should be logged. One from the gRPC call and one from the GraphQL endpoint.

## Notes

- The _React_ portion is created using the [create-react-app](https://github.com/facebook/create-react-app) tool.

- [gRPC](https://github.com/grpc/grpc) server is listening on port _9090_.

- The [gen-proto](gen-proto) file is a script to generate the protobuf files for _go_ and _typescript/js_.
