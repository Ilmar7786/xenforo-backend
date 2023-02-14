package v1

import (
	"context"
	"electronic_diary/app/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, ctx context.Context, userUC user.UseCase) {

	h := handler.Group("/v1")
	{
		newUserRouters(h, ctx, userUC)
	}
}
