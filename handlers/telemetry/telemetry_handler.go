package telemetry

import (
	"context"
	"log"

	pb "grpc-example/gen/telemetry/v1"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

type telemetryHandler struct {
	pb.UnimplementedTelemetryServiceServer
}

func NewTelemetryHandler() *telemetryHandler {
	return &telemetryHandler{}
}

// CheckTelemetry is the RPC method defined in your proto file.
func (s *telemetryHandler) Telemetry(ctx context.Context, req *pb.TelemetryRequest) (*pb.TelemetryResponse, error) {
	if req == nil {
		return nil, nil
	}

	// Setup MQTT client
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return nil, token.Error()
	}

	// Serialize req to JSON or any other format and publish
	resp := &pb.TelemetryResponse{
		Status: "OK",
		Status2: "OK2",
	}
	respBytes, err := proto.Marshal(resp)
	if err != nil {
		log.Println("Failed to serialize response:", err)
		return nil, err
	}
	token := client.Publish("telemetry", 0, false, respBytes)

	token.Wait()

	log.Println("Telemetry check OK!")
	return resp, nil
}