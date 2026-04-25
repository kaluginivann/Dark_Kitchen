package users

import (
	"context"

	"github.com/kaluginivann/Dark_Kitchen/models"
	"github.com/kaluginivann/Dark_Kitchen/pkg/db"
	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
)

type IRepository interface {
	CreateUser(ctx context.Context, username, email, password string) (*models.User, error)
}

type UserRepository struct {
	DB     *db.Database
	Logger logger.Interface
}

func NewRepository(db *db.Database, logger logger.Interface) *UserRepository {
	return &UserRepository{
		DB:     db,
		Logger: logger,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, username, email, password string) (*models.User, error) {
	query := `
		INSERT INTO users (
			username,
			email,
			password
		)
		VALUES ($1, $2, $3)
		RETURNING id, username, email
	`
	var user models.User

	if err := u.DB.Db.QueryRow(
		ctx,
		query,
		username,
		email,
		password,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
	); err != nil {
		return nil, err
	}
	return &user, nil
}
