import React from "react";
import { HelloRequest } from "./protos/hello/hello_pb";
import { GreeterClient } from "./protos/hello/HelloServiceClientPb";
import { QueryRequest } from "./protos/dgraph/dgraph_pb";
import { QueryClient } from "./protos/dgraph/DgraphServiceClientPb";

function App() {
  const envoyProxy = "http://localhost:8081";
  const greeter = new GreeterClient(envoyProxy);
  const request = new HelloRequest();
  request.setName("There!");
  greeter.sayHello(request, {}, (err, response) => {
    console.log(response.getReply());
  });

  // query dgraph using grpc. NOTE: syntax is GraphQL+-
  const dgraph = new QueryClient(envoyProxy);
  const grpcQuery = new QueryRequest();
  grpcQuery.setQuery(`{
    customer(func: has(Customer.username)){
      Customer.username
    }
  }`);
  dgraph.query(grpcQuery, {}, (err, response) => {
    console.log(JSON.parse(response.getResponse()));
  });

  // query dgraph using its graphql endpoint. NOTE: syntax is GraphQL
  const graphqlQuery = `{
  queryCustomer(filter: {username: {regexp: "/Mich.*/"}}) {
    username
  }
}`;
  fetch(`http://localhost:8080/graphql?query=${graphqlQuery}`)
    .then((resp) => resp.json())
    .then((data) => console.log(data));

  return <div className="App">Hello</div>;
}

export default App;
