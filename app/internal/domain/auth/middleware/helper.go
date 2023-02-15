package middleware

import (
	"net/http"
	"strings"

	"xenforo/app/pkg/api/jwt"

	"github.com/gin-gonic/gin"
)

func (i *Init) parseToken(c *gin.Context, privateKey string) (interface{}, error) {
	var accessToken string
	cookie, err := c.Cookie("access_token")

	authorizationHeader := c.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else if err == nil {
		accessToken = cookie
	}

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
		return nil, err
	}

	sub, err := jwt.ValidateToken(accessToken, privateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return nil, err
	}

	return sub, nil
}
