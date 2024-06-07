package main

import (
	"github.com/mynreden/microservices-go/common/database"
	"github.com/mynreden/microservices-go/common/utils"
	"github.com/mynreden/microservices-go/users/internal/services"
	"log"
	"net"

	"github.com/mynreden/microservices-go/common/api/proto"
	"google.golang.org/grpc"
)

func main() {
	address := utils.EnvString("USER_ADDR", "localhost:50051")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := database.GetDB(utils.EnvString("DSN", "postgresql://localhost:5432/finalAdvProg?user=postgres&password=sultan2004"))
	grpcServer := grpc.NewServer()

	proto.RegisterUserServiceServer(grpcServer, services.NewUserServiceServer(db))
	log.Printf("gRPC server started on %s...", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
