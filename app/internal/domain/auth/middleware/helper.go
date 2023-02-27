package middleware

import (
	"errors"
	"net/http"
	"strings"
	"xenforo/app/pkg/api"

	"github.com/gin-gonic/gin"
)

func (i *Init) parseToken(c *gin.Context, privateKey string) (interface{}, error) {
	var accessToken string

	authorizationHeader := c.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not logged in"})
		return nil, errors.New("access token empty")
	}

	sub, err := api.ValidateToken(accessToken, privateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return nil, err
	}

	return sub, nil
}
