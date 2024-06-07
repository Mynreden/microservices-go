package services

import (
	"context"
	"fmt"
	"github.com/mynreden/microservices-go/common/api/proto"
	"github.com/mynreden/microservices-go/common/models"
	"github.com/mynreden/microservices-go/posts/internal/repository"
)

type PostServiceServer struct {
	proto.UnimplementedPostServiceServer
	postRepository *repository.PostRepository
}

func NewPostServiceServer(postRepository *repository.PostRepository) *PostServiceServer {
	return &PostServiceServer{postRepository: postRepository}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.CreatePostResponse, error) {
	post := &models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserId:  req.UserId,
	}

	if err := s.postRepository.CreatePost(ctx, post); err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return &proto.CreatePostResponse{
		Id:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserId,
	}, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *proto.GetPostRequest) (*proto.GetPostResponse, error) {
	post, err := s.postRepository.GetPostByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return &proto.GetPostResponse{
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserId,
		Id:      post.ID,
	}, nil
}

func (s *PostServiceServer) GetPostsByUserId(ctx context.Context, req *proto.GetPostsByUserIdRequest) (*proto.GetPostsByUserIdResponse, error) {
	posts, err := s.postRepository.GetPostsByUserID(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts by user ID: %w", err)
	}

	var responsePosts []*proto.Post
	for _, post := range posts {
		responsePost := &proto.Post{
			Id:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserId:  post.UserId,
		}
		responsePosts = append(responsePosts, responsePost)
	}

	return &proto.GetPostsByUserIdResponse{
		Posts: responsePosts,
	}, nil
}

func (s *PostServiceServer) DeletePost(ctx context.Context, req *proto.DeletePostRequest) (*proto.DeletePostResponse, error) {
	if err := s.postRepository.DeletePostByID(ctx, req.Id); err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &proto.DeletePostResponse{Message: "Deleted successfully"}, nil
}

func (s *PostServiceServer) mustEmbedUnimplementedPostServiceServer() {}
