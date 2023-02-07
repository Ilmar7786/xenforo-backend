package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
