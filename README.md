# MQTT Telemetry Service - Go Examples

This repository provides an example of using MQTT protocol with Protocol Buffers (Protobuf) in Go for a simple telemetry service. 

The codebase demonstrates how to set up a MQTT client in Go, how to publish messages to an MQTT broker and how to subscribe to messages from a specific topic. It also provides examples of clients in TypeScript and JavaScript that can interact with the Go service.

**There's another example in this repo for more traditional gRPC and Protobuf communication - see the `main` branch**

## Structure

The repository is structured as follows:

- `mqtt-server-go/proto/telemetry/v1/telemetry.proto`: This is the protocol buffer definition for the telemetry service.
- `mqtt-server-go/main.go`: This file sets up a simple MQTT publisher in Go that sends messages to a specific topic at a given interval.
- `mqtt-client-go/main.go`: This file sets up a MQTT subscriber in Go that listens to messages from a specific topic.
- `mqtt-server-go/publish/publish.go`: This helper file contains the `PublishMessage` function which is used to send messages to the MQTT broker.
- `mqtt-server-go/utils/utils.go`: This file contains utility functions used in the project.

## Getting Started

Before running the examples, ensure that you have Go, Node.js, `buf`, and a running MQTT broker installed on your machine. 

For running the Go MQTT publisher:

1. Navigate to the `mqtt-server-go` directory.
2. Run `go run main.go` to start the MQTT publisher.

For running the Go MQTT subscriber:

1. Navigate to the `mqtt-client-go` directory.
2. Run `go run main.go` to start the MQTT subscriber.

In each case, the Go publisher will send a telemetry message to the "telemetry" topic at regular intervals, and the Go subscriber will listen to the "telemetry" topic and log the received messages.

This way, you can see a working example of a telemetry service using MQTT and Protobuf in Go, TypeScript, and JavaScript. If you update the `telemetry.proto` file and generate the Go code again, you can see how the changes propagate through the service.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is open source under the MIT license. See the [LICENSE](LICENSE) file for details.