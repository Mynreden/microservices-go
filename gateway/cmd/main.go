package main

import (
	"github.com/mynreden/microservices-go/common/api/proto"
	"github.com/mynreden/microservices-go/gateway/internal/app"
	configs "github.com/mynreden/microservices-go/gateway/internal/config"
	"github.com/mynreden/microservices-go/gateway/internal/handlers"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("wait a minute...")

	cfg, err := configs.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	userConn, err := grpc.Dial(cfg.UsersAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return
	}
	defer func(userConn *grpc.ClientConn) {
		err := userConn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(userConn)
	log.Println("Connected to UserService")
	userService := proto.NewUserServiceClient(userConn)

	postConn, err := grpc.Dial(cfg.PostsAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return
	}
	defer func(postConn *grpc.ClientConn) {
		err := postConn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(postConn)
	log.Println("Connected to UserService")
	postService := proto.NewPostServiceClient(postConn)

	handler := handlers.NewHandler(userService, postService)

	err = app.Server(cfg, handler.Routes())

	if err != nil {
		log.Println("Ooopss...\n", err)
		return
	}
}
