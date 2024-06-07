package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mynreden/microservices-go/common/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, `
		INSERT INTO users (username, email, password) 
		VALUES ($1, $2, $3) 
		RETURNING id`,
		user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(ctx, `
		SELECT id, username, email, password 
		FROM users 
		WHERE id = $1`, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUserByID(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `
		DELETE 
		FROM users 
		WHERE id = $1`, id)
	return err
}
