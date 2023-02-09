package http

import (
	userUC "electronic_diary/app/internal/domain/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitControllers(router *gin.Engine, db *gorm.DB) {
	userUseCase := userUC.NewUserUseCase(db)

	api := router.Group("/api")

	usersGroup := api.Group("/users")
	NewUserController(userUseCase).Register(usersGroup)
}
