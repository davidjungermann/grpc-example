package main

import (
	"time"

	"mqtt-server-go/publish"
	"mqtt-server-go/utils"
)

func main() {
	client := utils.SetUpMQTTClient()
	// Publish messages to the "telemetry" topic every 1 second
	go publish.Publish(client, utils.MQTT_TOPIC, time.Second)

	// Block forever, to keep the publisher running.
	select {}
}
