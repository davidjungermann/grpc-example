package publish

import (
	"log"
	"time"

	pb "mqtt-server-go/gen/telemetry/v1"
	"mqtt-server-go/rpi"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Publish(client mqtt.Client, topic string, interval time.Duration) {
	ticker := time.NewTicker(interval)

	// Loop and publish a message at each tick
	for range ticker.C {
		temp, err := rpi.ReadCPUTemp()

		if err != nil {
			log.Printf("Couldn't read temperature, assigning default value: %v", err)
			temp = 1337.0
		}

		telemetry := &pb.Telemetry{
			Temperature: temp,
			Timestamp:   timestamppb.New(time.Now().UTC()),
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
