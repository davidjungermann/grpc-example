package main

import (
	"log"

	pb "grpc-example/gen/telemetry/v1"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_subscriber")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return
	}

	if token := client.Subscribe("telemetry", 0, func(client mqtt.Client, msg mqtt.Message) {
		telemetryResponse := &pb.TelemetryResponse{}
		if err := proto.Unmarshal(msg.Payload(), telemetryResponse); err != nil {
			log.Println("Failed to deserialize message:", err)
		} else {
			log.Printf("Received message: %+v\n", telemetryResponse)
		}
	}); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// Block forever, to keep the subscriber running.
	select {}
}