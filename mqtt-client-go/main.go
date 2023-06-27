package main

import (
	"crypto/tls"
	"log"

	pb "buf.build/gen/go/djungermann/grpc-example/protocolbuffers/go/proto/telemetry/v1"

	"mqtt-server-go/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"
)

func main() {
	godotenv.Load()

	brokerMqttHost, err := utils.GetEnvironmentVariable("BROKER_MQTT_HOST")

	if err != nil {
		log.Fatalf("Failed to get MQTT host: %v", err)
	}

	brokerMqttPort, err := utils.GetEnvironmentVariable("BROKER_MQTT_PORT")

	if err != nil {
		log.Fatalf("Failed to get MQTT port: %v", err)
	}

	mqttUsername, err := utils.GetEnvironmentVariable("MQTT_USERNAME")

	if err != nil {
		log.Fatalf("Failed to get MQTT username: %v", err)
	}

	mqttPassword, err := utils.GetEnvironmentVariable("MQTT_PASSWORD")

	if err != nil {
		log.Fatalf("Failed to get MQTT password: %v", err)
	}

	// Create a TLS config that verifies the server certificate
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ClientAuth:         tls.NoClientCert,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker("ssl://"+brokerMqttHost+":"+brokerMqttPort)
	opts.SetClientID("hackweek-client")
	opts.SetTLSConfig(tlsConfig)
	opts.SetUsername(mqttUsername)
	opts.SetPassword(mqttPassword)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}

	// Subscribe to the "telemetry" topic
	if token := client.Subscribe("hackweek/telemetry", 0, func(client mqtt.Client, msg mqtt.Message) {
		telemetry := &pb.Telemetry{}
		if err := proto.Unmarshal(msg.Payload(), telemetry); err != nil {
			log.Println("Failed to deserialize message:", err)
		} else {
			log.Printf("Received message: %+v\n", telemetry)
		}
	}); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// Block forever, to keep the subscriber running.
	select {}
}
