package health

import (
	"context"
	"log"

	pb "grpc-example/gen/health/v1"
)

type healthHandler struct {
	pb.UnimplementedHealthServiceServer
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

// CheckHealth is the RPC method defined in your proto file.
func (s *healthHandler) CheckHealth(ctx context.Context, req *pb.CheckHealthRequest) (*pb.CheckHealthResponse, error) {
	if req == nil {
		return nil, nil
	}

	resp := &pb.CheckHealthResponse{
		Status: "OK",
	}
	log.Println("Health check OK!")
	return resp, nil
}
