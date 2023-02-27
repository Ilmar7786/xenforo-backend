package main

import (
	"context"

	"xenforo/app/internal/app"
	"xenforo/app/internal/config"
	"xenforo/app/pkg/logging"
)

// @title           Xenforo API
// @version         1.0
// @description     API Documentation Xenforo Server.

// @contact.name   API Support
// @contact.url    https://t.me/ilya112
// @contact.email  ilmar7786@yandex.ru

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.Info(ctx, "config initializing")
	cfg := config.GetConfig(ctx)

	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		logging.Fatal(ctx, err)
	}

	logging.Info(ctx, "Running Application")
	a.Run(ctx)
}
