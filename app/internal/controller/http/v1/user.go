package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/config"
	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
	"xenforo/app/pkg/api/jwt"
	"xenforo/app/pkg/api/validate"

	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	ctx    context.Context
	userUC user.UseCase
}

func newUserRouters(handler *gin.RouterGroup, ctx context.Context, authMiddleware middleware.Init, userUC user.UseCase) {
	r := userRoutes{
		ctx:    ctx,
		userUC: userUC,
	}

	h := handler.Group("/users")
	{
		h.POST("/sign-up", r.SignUp)
		h.POST("/sign-in", r.SignIn)
		h.GET("/info", authMiddleware.Auth(), r.UserInfo)
		h.PUT("/profile", authMiddleware.Auth(), r.UpdateProfile)
	}
}

func (r *userRoutes) SignIn(c *gin.Context) {
	cfg := config.GetConfig(r.ctx)
	body, err := validate.ParseAndValidateJSON[dto.UserAuthorizationDTO](c)
	if err != nil {
		return
	}

	currentUser, err := r.userUC.Authorization(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	// Generate Tokens
	accessToken, err := jwt.CreateToken(cfg.App.Jwt.AccessTokenExpiredIn, currentUser.ID, cfg.App.Jwt.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refreshToken, err := jwt.CreateToken(cfg.App.Jwt.RefreshTokenExpiredIn, currentUser.ID, cfg.App.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, cfg.App.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, cfg.App.Jwt.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", cfg.App.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"message": "validated",
		"tokens": gin.H{
			"access":  accessToken,
			"refresh": refreshToken,
		},
		"user": currentUser,
	})
}

func (r *userRoutes) SignUp(c *gin.Context) {
	body, err := validate.ParseAndValidateJSON[dto.UserCreateDTO](c)
	if err != nil {
		return
	}

	newUser, err := r.userUC.Create(body)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newUser)

}

func (r *userRoutes) UpdateProfile(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)
	c.JSON(http.StatusOK, currentUser)
}

func (r *userRoutes) UserInfo(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)
	c.JSON(http.StatusOK, currentUser)
}
