import React from "react";
import { HelloRequest } from "./protos/hello/hello_pb";
import { GreeterClient } from "./protos/hello/HelloServiceClientPb";

function App() {
  const client = new GreeterClient("http://localhost:8081");

  const request = new HelloRequest();
  request.setName("There!");

  client.sayHello(request, {}, (err, response) => {
    console.log(response);
    console.log(response.getReply());
    console.log(err);
  });
  return <div className="App">Hello</div>;
}

export default App;
