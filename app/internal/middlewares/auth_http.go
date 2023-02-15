package middlewares

import (
	"context"
	"electronic_diary/app/internal/config"
	"electronic_diary/app/internal/domain/user"
	"electronic_diary/app/pkg/api/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(ctx context.Context, userUC user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			return
		}

		cfg := config.GetConfig(ctx)
		sub, err := jwt.ValidateToken(accessToken, cfg.App.Jwt.AccessTokenPrivateKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		currentUser, err := userUC.FindByID(fmt.Sprint(sub))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}

		c.Set("currentUser", *currentUser)
		c.Next()
	}
}
