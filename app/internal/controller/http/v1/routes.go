package v1

import (
	"context"
	"xenforo/app/internal/domain/list_lock"

	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type UseCases struct {
	UserUC     user.UseCase
	ListLockUC list_lock.UseCase
}

func NewRouter(handler *gin.RouterGroup, ctx context.Context, authMiddleware middleware.Init, useCases UseCases) {
	v1 := handler.Group("/v1")

	newUserRouters(v1, ctx, authMiddleware, useCases.UserUC)

	private := v1.Group("/admin")
	newListLockRouters(private, ctx, authMiddleware, useCases.ListLockUC)

}
