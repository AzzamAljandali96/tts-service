package main

import (
	"log"
	"net"

	"github.com/AzzamAljandali96/tts-service/internal/config"
	"github.com/AzzamAljandali96/tts-service/internal/transport/grpc"
)

func main() {
	cfg := config.Load()

	lis, err := net.Listen("tcp", cfg.GRPCAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(cfg)
	log.Printf("Starting gRPC server on %s", cfg.GRPCAddress)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
