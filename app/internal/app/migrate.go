package app

import (
	"context"
	"electronic_diary/app/internal/config"
	"electronic_diary/app/internal/domain/user/model"
	"electronic_diary/app/pkg/client/gorm_postgesql"
	"electronic_diary/app/pkg/logging"
)

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	cfg := config.GetConfig(ctx)

	pgConfig := gorm_postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database,
	)
	pgClient := gorm_postgesql.NewClient(pgConfig)

	logging.Info(ctx, "start migrations")
	err := pgClient.AutoMigrate(&model.User{})

	if err != nil {
		logging.Error(ctx, err)
	}

	logging.Info(ctx, "migration was successful")
}
