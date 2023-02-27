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

// @Summary Бан
// @Security ApiKeyAuth
// @Tags admin
// @Description Блокировка и разблокировка пользователя
// @Accept json
// @Produce json
// @Param input body dto.UserBanDTO true "credentials"
// @Param user_id path string true "ID пользователя"
// @Success 200 {boolean} true
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /admin/users/{user_id}/ban [put]
func (a *adminHandler) userBan(c *gin.Context) {
	userID := c.Param("id")

	var input dto.UserBanDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userID
	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userBanned, err := a.userUC.BanUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, userBanned)
}

// @Summary Список пользователей
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /admin/users [get]
func (a *adminHandler) findUsers(c *gin.Context) {
	users := a.userUC.FindAll()

	c.JSON(http.StatusOK, users)
}
