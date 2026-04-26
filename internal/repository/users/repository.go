package users

import (
	"context"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
	"github.com/kaluginivann/Dark_Kitchen/internal/models"
	"github.com/kaluginivann/Dark_Kitchen/internal/repository/sql"
)

type Repository struct {
	db     sql.DBExecutor
	logger *zap.Logger
}

func NewUsersRepository(db sql.DBExecutor, logger *zap.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) Create(ctx context.Context, user *models.User) (*int, error) {
	var userID int

	r.logger.Debug("repository.users.Create")

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	if err := r.db.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&userID); err != nil {
		r.logger.Error("repository.Users.Create failed", zap.Error(err))
		return nil, err
	}

	r.logger.Debug("repository.Users.Create success", zap.Int("created_user_id", userID))

	return &userID, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*models.User, error) {

	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = $1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		r.logger.Error("repository.Users.GetById failed to query", zap.Error(err))
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		r.logger.Error("repository.Users.GetById failed to collect user", zap.Error(err))
	}

	return &user, nil
}
