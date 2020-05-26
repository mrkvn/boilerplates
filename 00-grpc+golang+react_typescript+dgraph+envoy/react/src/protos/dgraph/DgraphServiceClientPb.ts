/**
 * @fileoverview gRPC-Web generated client stub for dgraph
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  QueryRequest,
  QueryResponse} from './dgraph_pb';

export class QueryClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoQuery = new grpcWeb.AbstractClientBase.MethodInfo(
    QueryResponse,
    (request: QueryRequest) => {
      return request.serializeBinary();
    },
    QueryResponse.deserializeBinary
  );

  query(
    request: QueryRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: QueryResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/dgraph.Query/Query',
      request,
      metadata || {},
      this.methodInfoQuery,
      callback);
  }

}

