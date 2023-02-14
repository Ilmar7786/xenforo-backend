package v1

import (
	"context"
	"electronic_diary/app/internal/config"
	"electronic_diary/app/internal/domain/user"
	"electronic_diary/app/internal/domain/user/dto"
	"electronic_diary/app/internal/domain/user/model"
	"electronic_diary/app/internal/middlewares"
	"electronic_diary/app/pkg/api/jwt"
	"electronic_diary/app/pkg/api/validate"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	ctx    context.Context
	userUC user.UseCase
}

func newUserRouters(handler *gin.RouterGroup, ctx context.Context, userUC user.UseCase) {
	r := userRoutes{
		ctx:    ctx,
		userUC: userUC,
	}

	h := handler.Group("/users")
	{
		h.POST("/sign-up", r.SignUp)
		h.POST("/sign-in", r.SignIn)
		h.GET("/info", r.UserInfo)
		h.PUT("/profile", middlewares.Auth(ctx, userUC), r.UpdateProfile)
	}
}

func (r *userRoutes) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "hello from controller",
	})
}

func (r *userRoutes) SignUp(c *gin.Context) {
	cfg := config.GetConfig(r.ctx)
	body, err := validate.ParseAndValidateJSON[dto.UserAuthorizationDTO](c)
	if err != nil {
		return
	}

	newUser, err := r.userUC.Authorization(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	// Generate Tokens
	accessToken, err := jwt.CreateToken(cfg.App.Jwt.AccessTokenExpiredIn, newUser.ID, cfg.App.Jwt.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refreshToken, err := jwt.CreateToken(cfg.App.Jwt.RefreshTokenExpiredIn, newUser.ID, cfg.App.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, cfg.App.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, cfg.App.Jwt.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", cfg.App.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, false)

	input, err := validate.ParseAndValidateJSON[dto.UserCreateDTO](c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated", "user": input})
}

func (r *userRoutes) UpdateProfile(c *gin.Context) {

}

func (r *userRoutes) UserInfo(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)
	c.JSON(http.StatusOK, currentUser)
}
