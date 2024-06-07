package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpcServer *grpc.Server
	address    string
}

func NewGRPCServer(address string) *GRPCServer {
	return &GRPCServer{
		grpcServer: grpc.NewServer(),
		address:    address,
	}
}

func (s *GRPCServer) Start() {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Failed to listen on %v: %v", s.address, err)
	}

	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *GRPCServer) Stop() {
	s.grpcServer.GracefulStop()
}

func (s *GRPCServer) GetServer() *grpc.Server {
	return s.grpcServer
}
