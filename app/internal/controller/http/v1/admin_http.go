package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"

	"github.com/gin-gonic/gin"
)

type adminRoutes struct {
	ctx    context.Context
	userUC user.UseCase
}

func newAdminRoutes(handler *gin.RouterGroup, ctx context.Context, middleware middleware.Init, userUC user.UseCase) {
	r := adminRoutes{
		ctx:    ctx,
		userUC: userUC,
	}

	users := handler.Group("/users")
	{
		users.PUT("/:id/ban", r.UserBan)
	}
}

func (r *adminRoutes) UserBan(c *gin.Context) {
	userID := c.Param("id")

	var input dto.UserBanDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userID
	if err := input.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userBanned, err := r.userUC.BanUser(input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, userBanned)
}
