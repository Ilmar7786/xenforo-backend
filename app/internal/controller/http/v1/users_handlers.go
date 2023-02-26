package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/config"
	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
	"xenforo/app/pkg/api/jwt"
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

func (r *usersHandler) signIn(c *gin.Context) {
	cfg := config.GetConfig(r.ctx)
	input, err := validate.ParseAndValidateJSON[dto.UserAuthorizationDTO](c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	currentUser, err := r.userUC.Authorization(input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
	}

	// Generate Tokens
	accessToken, err := jwt.CreateToken(
		cfg.App.Jwt.AccessTokenExpiredIn, currentUser.ID,
		cfg.App.Jwt.AccessTokenPrivateKey,
	)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	refreshToken, err := jwt.CreateToken(
		cfg.App.Jwt.RefreshTokenExpiredIn, currentUser.ID,
		cfg.App.Jwt.RefreshTokenPrivateKey,
	)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": currentUser,
		"tokens": gin.H{
			"access":  accessToken,
			"refresh": refreshToken,
		},
	})
}

func (r *usersHandler) signUp(c *gin.Context) {
	input, err := validate.ParseAndValidateJSON[dto.UserCreateDTO](c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	newUser, err := r.userUC.Registration(input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (r *usersHandler) updateProfile(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)
	c.JSON(http.StatusOK, currentUser)
}

func (r *usersHandler) userInfo(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)
	c.JSON(http.StatusOK, currentUser)
}
