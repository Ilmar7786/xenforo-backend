package http

import (
	"electronic_diary/app/internal/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	userUC user.UseCase
}

func NewUserHandlers(userUC user.UseCase) user.Handlers {
	return userHandlers{userUC: userUC}
}

func (t userHandlers) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create",
	})
}

func (t userHandlers) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update",
	})
}

func (t userHandlers) FindByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "FindByID",
	})
}

func (t userHandlers) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete",
	})
}

func (t userHandlers) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "FindAll",
	})
}
