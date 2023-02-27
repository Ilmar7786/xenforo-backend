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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not enough rights"})
			return
		}

		if !currentUser.IsAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access is denied"})
			return
		}

		c.Set("user", *currentUser)
		c.Next()
	}
}
