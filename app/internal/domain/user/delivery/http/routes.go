package http

import (
	"electronic_diary/app/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(teacherGroup *gin.RouterGroup, h user.Handlers) {
	teacherGroup.GET("/", h.FindAll)
	teacherGroup.GET("/:id", h.FindByID)
	teacherGroup.POST("/", h.Create)
	teacherGroup.PATCH("/", h.Update)
	teacherGroup.DELETE("/", h.Delete)
}
