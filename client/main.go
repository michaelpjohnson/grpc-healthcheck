package main

import (
	"log"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	h := healthpb.NewHealthClient(conn)
	resp, err := h.Check(context.Background(), &healthpb.HealthCheckRequest{})
	log.Println(resp)

}
