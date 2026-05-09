package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaluginivann/Dark_Kitchen/config"
	"github.com/kaluginivann/Dark_Kitchen/internal/models"
	"github.com/kaluginivann/Dark_Kitchen/internal/repository/sql"

	"go.uber.org/fx"
)

var Module = fx.Module("repository",
	fx.Provide(newPool),
	fx.Provide(newDBExecutor),
	fx.Provide(New),
	fx.Provide(
		provideUsersRepository,
	),
	fx.Invoke(manageDBLifecycle),
)

func newPool(cfg *config.Config) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.Postgres.DSN)
}

func newDBExecutor(pool *pgxpool.Pool) sql.DBExecutor {
	return pool
}

func manageDBLifecycle(lc fx.Lifecycle, repo Repository) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return repo.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()

			return repo.Stop(ctx)
		},
	})
}

func provideUsersRepository(repo Repository) models.UsersRepository {
	return repo.UsersRepository()
}
