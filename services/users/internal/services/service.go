package services

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mynreden/microservices-go/common/api/proto"
	"github.com/mynreden/microservices-go/common/models"
	"github.com/mynreden/microservices-go/users/internal/repository"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	repo *repository.UserRepository
}

func NewUserServiceServer(db *pgxpool.Pool) *UserServiceServer {
	userRepo := repository.NewUserRepository(db)
	return &UserServiceServer{
		repo: userRepo,
	}
}

func (u *UserServiceServer) CreateUser(ctx context.Context, request *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password, // In a real app, hash the password before storing it
	}

	id, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{
		Id:       id,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Message:  "User created successfully",
	}, nil
}

func (u *UserServiceServer) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := u.repo.GetUserByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Message:  "User received successfully",
	}, nil
}

func (u *UserServiceServer) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := u.repo.DeleteUserByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}

func (u *UserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	// This method is to ensure forward compatibility
}
