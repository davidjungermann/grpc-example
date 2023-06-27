package main

import (
	"log"
	"time"

	pb "mqtt-server-go/gen/telemetry/v1"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

// publish sends a message to an MQTT bus at a given interval
func publish(client mqtt.Client, topic string, interval time.Duration) {
	ticker := time.NewTicker(interval)

	// Loop and publish a message at each tick
	for range ticker.C {
		telemetry := &pb.Telemetry{
			Status : "OK",
			// Add other fields as needed
		}

		data, err := proto.Marshal(telemetry)
		if err != nil {
			log.Printf("Error marshalling to proto: %v", err)
			continue
		}

		token := client.Publish(topic, 0, false, data)
		if token.Error() != nil {
			log.Printf("Error publishing to topic %s: %v", topic, token.Error())
			continue
		}

		log.Printf("Published message to topic %s", topic)
	}
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_publisher")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	// Publish messages to the "telemetry" topic every 1 second
	go publish(client, "telemetry", time.Second)

	// Block forever, to keep the publisher running.
	select {}
}
