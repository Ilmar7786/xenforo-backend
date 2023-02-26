package app

import (
	"context"

	"xenforo/app/internal/config"
	MailModel "xenforo/app/internal/domain/mail/model"
	UserModel "xenforo/app/internal/domain/user/model"
	"xenforo/app/pkg/client/gorm_postgesql"
	"xenforo/app/pkg/logging"
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
	err := pgClient.AutoMigrate(&UserModel.User{}, MailModel.MailActivate{})

	if err != nil {
		logging.Error(ctx, err)
	}

	logging.Info(ctx, "migration was successful")
}
