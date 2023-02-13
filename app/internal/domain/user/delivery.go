package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	Registration(ctx *gin.Context)
	Authorization(ctx *gin.Context)
}
