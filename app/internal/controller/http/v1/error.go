package v1

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Error interface{} `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg interface{}) {
	c.AbortWithStatusJSON(code, response{msg})
}
