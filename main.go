package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb_health "grpc-example/gen/health/v1"
	health_handler "grpc-example/handlers"
	"grpc-example/utils"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	const defaultGRPCPort = "50051"
	const defaultHTTPPort = "8080"

	grpcAddress := utils.GetPortFromEnv("GRPC_PORT", defaultGRPCPort)
	httpAddress := utils.GetPortFromEnv("HTTP_PORT", defaultHTTPPort)

	// Initialize handlers
	healthHandler := health_handler.NewHealthHandler()

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach handler to the server
	pb_health.RegisterHealthServiceServer(s, healthHandler)

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Serve gRPC Server
	log.Printf("Serving gRPC on %s", grpcAddress)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// Register GW handlers for reverse proxy
	err = pb_health.RegisterHealthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    httpAddress,
		Handler: gwmux,
	}

	log.Printf("Serving HTTP gRPC-Gateway on %s", httpAddress)
	log.Fatalln(gwServer.ListenAndServe())
}
