package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
	"xenforo/app/pkg/api/validate"

	"github.com/gin-gonic/gin"
)

type usersHandler struct {
	ctx    context.Context
	userUC user.UseCase
}

func newUserHandler(ctx context.Context, userUC user.UseCase) *usersHandler {
	return &usersHandler{
		ctx:    ctx,
		userUC: userUC,
	}
}

// @Summary Авторизация
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.UserAuthorizationDTO true "credentials"
// @Success 200 {object} model.UserAndTokens
// @Failure 400 {object} errorResponse
// @Router /users/sign-in [post]
func (r *usersHandler) signIn(c *gin.Context) {
	input, err := validate.ParseAndValidateJSON[dto.UserAuthorizationDTO](c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	currentUser, err := r.userUC.Authorization(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, currentUser)
}

// @Summary Регистрация
// @Tags users
// @Description <b>RedirectActiveEmail</b> - ссылка редиректа для активации почты. Эта ссылка указывается в письме
// @Accept json
// @Produce json
// @Param input body dto.UserRegistrationDTO true "credentials"
// @Success 200 {object} model.UserAndTokens
// @Failure 400 {object} errorResponse
// @Router /users/sign-up [post]
func (r *usersHandler) signUp(c *gin.Context) {
	input, err := validate.ParseAndValidateJSON[dto.UserRegistrationDTO](c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := r.userUC.Registration(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// @Summary Обновить
// @Security ApiKeyAuth
// @Tags users
// @Description Обновления атрибутов пользователя
// @Accept json
// @Produce json
// @Param input body dto.UserUpdateDTO true "credentials"
// @Success 200 {object} model.User
// @Failure 400 {object} errorResponse
// @Router /users/profile [put]
func (r *usersHandler) updateProfile(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)
	c.JSON(http.StatusOK, currentUser)
}

// @Summary Информация о пользователи
// @Security ApiKeyAuth
// @Tags users
// @Description Данные о пользователе
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /users/info [get]
func (r *usersHandler) userInfo(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)
	c.JSON(http.StatusOK, currentUser)
}
