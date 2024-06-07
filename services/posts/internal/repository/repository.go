package repository

import (
	"context"
	"fmt"
	"github.com/driftprogramming/pgxpoolmock"

	"github.com/mynreden/microservices-go/common/models"
)

type PostRepository struct {
	db pgxpoolmock.PgxPool
}

func NewPostRepository(db pgxpoolmock.PgxPool) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (title, content, userId)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, post.Title, post.Content, post.UserId).Scan(&post.ID)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (r *PostRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post

	query := `
		SELECT id, title, content, userid
		FROM posts
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserId)
	if err != nil {
		return nil, fmt.Errorf("error getting post by ID: %w", err)
	}

	return &post, nil
}

func (r *PostRepository) DeletePostByID(ctx context.Context, id string) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting post by ID: %w", err)
	}
	return nil
}

func (r *PostRepository) GetPostsByUserID(ctx context.Context, userID string) ([]*models.Post, error) {
	query := `
		SELECT id, title, content, userid
		FROM posts
		WHERE userid = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying posts by user ID: %w", err)
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserId); err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over post rows: %w", err)
	}

	return posts, nil
}
