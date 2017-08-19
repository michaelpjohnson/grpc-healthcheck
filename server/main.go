package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	gRPCServer *grpc.Server
)

type server struct{}

func (s *server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	healthpb.RegisterHealthServer(srv, &server{})

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
