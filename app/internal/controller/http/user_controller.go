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
	handlers := UserHttp.NewUserHandlers(userUC)

	return &UserController{
		userHandlers: handlers,
	}
}

func (c UserController) Register(group *gin.RouterGroup) {
	group.GET("/", c.userHandlers.FindAll)
	group.GET("/:id", c.userHandlers.FindByID)
	group.POST("/", c.userHandlers.Create)
	group.PATCH("/", c.userHandlers.Update)
	group.DELETE("/", c.userHandlers.Delete)
}
