package v1

import (
	"context"

	"xenforo/app/internal/config"
	"xenforo/app/internal/domain/event"
	"xenforo/docs"

	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/sport"
	"xenforo/app/internal/domain/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller struct {
	UserUC  user.UseCase
	SportUC sport.UseCase
	EventUC event.UseCase
}

const prefix = "/api"

func NewRouter(router *gin.Engine, ctx context.Context, middleware middleware.Init, controller Controller) {
	cfg := config.GetConfig(ctx)

	public := router.Group(prefix + "/v1")
	{
		if cfg.App.IsDebug {
			public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			docs.SwaggerInfo.BasePath = public.BasePath()
		}

		users := public.Group("/users")
		{
			h := newUserHandler(ctx, controller.UserUC)
			users.POST("/sign-up", h.signUp)
			users.POST("/sign-in", h.signIn)

			users.Use(middleware.Auth())
			users.GET("/info", h.userInfo)
			users.PUT("/profile", h.updateProfile)
		}

		sports := public.Group("/sports")
		{
			h := newSportHandler(ctx, controller.SportUC)
			sports.GET("/", h.GetList)
		}

		events := public.Group("/events")
		{
			h := newEventHandler(ctx, controller.EventUC)
			events.GET("/", h.GetList)
		}
	}

	private := public.Group("/admin")
	private.Use(middleware.AdminAuth())
	{
		h := newAdminHandler(ctx, controller.UserUC)
		users := private.Group("/users")
		{
			users.PUT("/:id/ban", h.userBan)
			users.GET("", h.findUsers)
		}
	}
}
