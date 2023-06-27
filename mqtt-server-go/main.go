package main

import (
	"crypto/tls"
	"log"
	"time"

	"mqtt-server-go/publish"
	"mqtt-server-go/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
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
	opts.SetClientID("hackweek-server")
	opts.SetTLSConfig(tlsConfig)
	opts.SetUsername(mqttUsername)
	opts.SetPassword(mqttPassword)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}

	// Publish messages to the "telemetry" topic every 1 second
	go publish.Publish(client, "hackweek/telemetry", time.Second)

	// Block forever, to keep the publisher running.
	select {}
}
