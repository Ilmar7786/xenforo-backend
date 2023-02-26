package v1

import (
	"context"

	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/sport"
	"xenforo/app/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type UseCases struct {
	UserUC  user.UseCase
	SportUC sport.UseCase
}

func NewRouter(handler *gin.RouterGroup, ctx context.Context, middleware middleware.Init, useCases UseCases) {
	public := handler.Group("/v1")
	{
		users := public.Group("/users")
		{
			h := newUserHandler(ctx, useCases.UserUC)
			users.POST("/sign-up", h.signUp)
			users.POST("/sign-in", h.signIn)

			users.Use(middleware.Auth())
			users.GET("/info", h.userInfo)
			users.PUT("/profile", h.updateProfile)
		}

		sports := public.Group("/sports")
		{
			h := newSportHandler(ctx, useCases.SportUC)
			sports.GET("/", h.GetList)
		}
	}

	private := public.Group("/admin")
	private.Use(middleware.AdminAuth())
	{
		h := newAdminHandler(ctx, useCases.UserUC)
		users := private.Group("/users")
		{
			users.PUT("/:id/ban", h.userBan)
		}
	}
}
