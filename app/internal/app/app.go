package app

import (
	"context"
	http2 "electronic_diary/app/internal/domain/user/delivery/http"
	"electronic_diary/app/internal/domain/user/usecase"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"electronic_diary/app/internal/config"
	"electronic_diary/app/pkg/client/gorm_postgesql"
	"electronic_diary/app/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type App struct {
	cfg *config.Config

	router     *gin.Engine
	httpServer *http.Server
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {
	logging.Info(ctx, "router initializing")
	if !cfg.App.IsDebug {
		if err := os.Setenv(gin.EnvGinMode, gin.ReleaseMode); err != nil {
			logging.Error(ctx, err)
		}
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	// Database postgresql
	pgConfig := gorm_postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database,
	)
	pgClient := gorm_postgesql.NewClient(pgConfig)

	// Init useCases
	userUC := usecase.NewUserUseCase(pgClient)

	// Init handlers
	handlers := http2.NewUserHandlers(userUC)

	// Init routing
	api := router.Group("/api")

	usersGroup := api.Group("/users")

	http2.MapUserRoutes(usersGroup, handlers)

	return App{
		cfg:    cfg,
		router: router,
	}, nil
}

func (a *App) Run(ctx context.Context) {
	a.startHTTP(ctx)
}

func (a *App) startHTTP(ctx context.Context) {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"IP":   a.cfg.HTTP.IP,
		"Port": a.cfg.HTTP.Port,
	})
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		logger.WithError(err).Fatal("failed to create listener")
	}

	logger.WithFields(map[string]interface{}{
		"AllowedMethods":     a.cfg.HTTP.CORS.AllowedMethods,
		"AllowedOrigins":     a.cfg.HTTP.CORS.AllowedOrigins,
		"AllowCredentials":   a.cfg.HTTP.CORS.AllowCredentials,
		"AllowedHeaders":     a.cfg.HTTP.CORS.AllowedHeaders,
		"OptionsPassthrough": a.cfg.HTTP.CORS.OptionsPassthrough,
		"ExposedHeaders":     a.cfg.HTTP.CORS.ExposedHeaders,
		"Debug":              a.cfg.HTTP.CORS.Debug,
	})
	c := cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warning("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
}
