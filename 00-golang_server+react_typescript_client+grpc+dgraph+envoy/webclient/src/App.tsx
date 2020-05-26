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
    console.log(response);
    console.log(response.getReply());
    console.log(err);
  });

  const dgraph = new QueryClient(envoyProxy);
  const query = new QueryRequest();
  query.setQuery(`{
    people(func: has(name)){
      name
      age
    }
  }`);
  dgraph.query(query, {}, (err, response) => {
    console.log(response);
    console.log(response.getResponse());
    console.log(err);
  });

  return <div className="App">Hello</div>;
}

export default App;
