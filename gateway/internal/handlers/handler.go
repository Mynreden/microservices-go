package handlers

import "github.com/mynreden/microservices-go/common/api/proto"

type Handler struct {
	userService proto.UserServiceClient
	postService proto.PostServiceClient
}

func NewHandler(userService proto.UserServiceClient, postService proto.PostServiceClient) *Handler {
	return &Handler{userService,
		postService}
}
