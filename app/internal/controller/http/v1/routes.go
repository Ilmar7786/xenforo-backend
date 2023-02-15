package v1

import (
	"context"

	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.RouterGroup, ctx context.Context, authMiddleware middleware.Init, userUC user.UseCase) {
	v1 := handler.Group("/v1")
	{
		newUserRouters(v1, ctx, authMiddleware, userUC)
	}
}
