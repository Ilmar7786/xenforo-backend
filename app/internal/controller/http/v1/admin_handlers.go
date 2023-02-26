package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	ctx    context.Context
	userUC user.UseCase
}

func newAdminHandler(ctx context.Context, userUC user.UseCase) *adminHandler {
	return &adminHandler{
		ctx:    ctx,
		userUC: userUC,
	}
}

func (r *adminHandler) userBan(c *gin.Context) {
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
