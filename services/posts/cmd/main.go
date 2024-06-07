package main

import (
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/mynreden/microservices-go/common/database"
	"github.com/mynreden/microservices-go/common/utils"
	"github.com/mynreden/microservices-go/posts/internal/repository"
	"github.com/mynreden/microservices-go/posts/internal/services"
	"log"
	"net"

	"github.com/mynreden/microservices-go/common/api/proto"
	"google.golang.org/grpc"
)

func main() {
	address := utils.EnvString("USER_ADDR", "localhost:50052")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var db pgxpoolmock.PgxPool = database.GetDB(utils.EnvString("DSN", "postgresql://localhost:5432/finalAdvProg?user=postgres&password=sultan2004"))
	grpcServer := grpc.NewServer()

	repo := repository.NewPostRepository(db)
	proto.RegisterPostServiceServer(grpcServer, services.NewPostServiceServer(repo))
	log.Printf("gRPC server started on %s...", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
