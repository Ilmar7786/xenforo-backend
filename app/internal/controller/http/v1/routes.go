package v1

import (
	"context"

	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type UseCases struct {
	UserUC user.UseCase
}

func NewRouter(handler *gin.RouterGroup, ctx context.Context, middleware middleware.Init, useCases UseCases) {
	public := handler.Group("/v1")
	private := public.Group("/admin")
	private.Use(middleware.AdminAuth())

	newAdminRoutes(private, ctx, middleware, useCases.UserUC)
	newUserRouters(public, ctx, middleware, useCases.UserUC)
}
