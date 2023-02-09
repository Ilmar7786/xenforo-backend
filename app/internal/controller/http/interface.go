package http

import "github.com/gin-gonic/gin"

type ControllerRegister interface {
	Register(route *gin.RouterGroup)
}
