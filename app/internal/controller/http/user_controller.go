package http

import (
	"electronic_diary/app/internal/domain/user"
	UserHttp "electronic_diary/app/internal/domain/user/delivery/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userHandlers user.Handlers
}

func NewUserController(userUC user.UseCase) ControllerRegister {
	handlers := UserHttp.NewHandlers(userUC)

	return &UserController{
		userHandlers: handlers,
	}
}

func (c UserController) Register(group *gin.RouterGroup) {
	group.POST("/sign-up", c.userHandlers.Registration)
	group.POST("/sign-in", c.userHandlers.Authorization)
}
