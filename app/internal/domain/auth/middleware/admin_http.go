package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (i *Init) AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, err := i.parseToken(c, i.privateKey)
		if err != nil {
			return
		}

		currentUser, err := i.userUC.FindByID(fmt.Sprint(sub))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Access is denied"})
			return
		}

		if !currentUser.IsAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}

		c.Set("user", *currentUser)
		c.Next()
	}
}
