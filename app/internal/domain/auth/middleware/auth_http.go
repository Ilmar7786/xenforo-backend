package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (i *Init) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, err := i.parseToken(c, i.privateKey)
		if err != nil {
			return
		}

		currentUser, err := i.userUC.FindByID(fmt.Sprint(sub))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "The user belonging to this token no logger exists"})
			return
		}

		if currentUser.IsBanned {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message, ": "The user is blocked"})
			return
		}

		c.Set("user", *currentUser)
		c.Next()
	}
}
