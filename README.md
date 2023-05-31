# gRPC Health Check - TypeScript and JavaScript Example

This repository provides an example of using gRPC with Protobuf in TypeScript and JavaScript for a simple health check service. 

The codebase demonstrates how to set up the gRPC server in Go, and how to create TypeScript and JavaScript clients that can interact with it.

## Structure

The repository is structured as follows:

- `proto/health/v1/health.proto`: This is the protocol buffer definition for the health check service.
- `client-server-ts`: This directory contains the TypeScript client and server code. 
  - `client.ts` and `server.ts` are the main files to look at for seeing how to use the generated code to create a gRPC server and client.
- `client-js`: Similar to the TypeScript directory, this contains a JavaScript client for the health check service.
  - `grpc.js` shows how to use the generated JavaScript code to create a gRPC client.
- `main.go` and `handlers/health_handler.go`: These files set up a simple gRPC server in Go using the protocol buffer definitions. 
- `Makefile`: This file provides convenient commands for building and running the project.

## Getting Started

Before running the examples, ensure that you have Node.js, Go, and `buf` installed on your machine. 

For running the Go server:

1. Run `make server` from the root directory to start the Go gRPC server.

For the TypeScript client:

1. Navigate to the `client-server-ts` directory.
2. Run `buf generate buf.build/djungermann/grpc-example`. This fetches the proto files from the remote registry.
3. Run `npm install` to install the necessary dependencies.
4. Run `npx tsx server.ts` to start the TypeScript server.
5. Make a request to a server using `npx tsx grpc.ts`
6. Make a request to a server using `npx tsx http.ts`

For the JavaScript client:

1. Navigate to the `client-js` directory.
2. Run `buf generate buf.build/djungermann/grpc-example`. This fetches the proto files from the remote registry.
3. Run `npm install` to install the necessary dependencies.
4. Run `node grpc.js` to make a request to a server via gRPC.
5. Run `node http.js` to make a request to a server via HTTP.

In each case, the client will send a health check request to the server and log the response.

This way, you can run the Go service, and make requests from both Typescript and Javascript, but you can also run the Typescript server, that has generated both client and server code and make the same request that way.
If you update the proto file and push that to Buf, you can try out how to propagate changes, and then implement them in the various examples. 

**It's also possible to use connect-es to make the HTTP request, but I'll leave that up to the reader**

The TypeScript and JavaScript clients demonstrate how to use the protocol buffer messages and service stubs generated from the `.proto` files to make requests to a gRPC server.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is open source under the MIT license. See the [LICENSE](LICENSE) file for details.
