package repository

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaluginivann/Dark_Kitchen/internal/models"
	"github.com/kaluginivann/Dark_Kitchen/internal/repository/sql"
	"github.com/kaluginivann/Dark_Kitchen/internal/repository/users"
)

type Repository interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error

	UsersRepository() models.UsersRepository
}

type repository struct {
	dbExecutor      sql.DBExecutor
	pool            *pgxpool.Pool
	logger          *zap.Logger
	usersRepository models.UsersRepository
}

func (r *repository) Start(ctx context.Context) error {
	r.logger.Info("Pinging Postgres DB")

	err := r.pool.Ping(ctx)
	if err != nil {
		r.logger.Error("Postgres DB is not available", zap.Error(err))
		return err
	}

	r.logger.Info("Postgres DB is ready")
	return nil
}

func (r *repository) Stop(ctx context.Context) error {
	startTime := time.Now()

	r.logger.Info("Closing Postgres DB connection pool...",
		zap.Int32("active_conns", r.pool.Stat().AcquiredConns()))

	done := make(chan struct{})

	go func() {
		r.pool.Close()
		close(done)
	}()

	select {
	case <-done:
		r.logger.Info("Postgres DB closed gracefully",
			zap.Duration("duration", time.Since(startTime)))
		return nil
	case <-ctx.Done():
		r.logger.Warn("Shutdown timeout, some connections may be interrupted",
			zap.Int32("remaining_conns", r.pool.Stat().AcquiredConns()),
			zap.Duration("waited", time.Since(startTime)))
		return ctx.Err()
	}
}

func (r *repository) UsersRepository() models.UsersRepository {
	return r.usersRepository
}

func New(pool *pgxpool.Pool, dbExecutor sql.DBExecutor, logger *zap.Logger) Repository {
	return &repository{
		pool:            pool,
		dbExecutor:      dbExecutor,
		logger:          logger,
		usersRepository: users.NewUsersRepository(pool, logger),
	}
}
