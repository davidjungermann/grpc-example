package publish

import (
	"log"
	"time"

	pb "mqtt-server-go/gen/telemetry/v1"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

func Publish(client mqtt.Client, topic string, interval time.Duration) {
	ticker := time.NewTicker(interval)

	// Loop and publish a message at each tick
	for range ticker.C {
		telemetry := &pb.Telemetry{
			Status: "OK",
			// Add other fields as needed
		}

		data, err := proto.Marshal(telemetry)
		if err != nil {
			log.Printf("Error marshalling to proto: %v", err)
			continue
		}

		publishMessage(client, topic, data, false)
	}
}

func publishMessage(client mqtt.Client, topic string, payload []byte, retain bool) {
	token := client.Publish(topic, 1, retain, payload)
	if token.Error() != nil {
		log.Printf("Error publishing to topic %s: %v", topic, token.Error())
		return
	}
	log.Printf("Published message to topic %s", topic)
}
