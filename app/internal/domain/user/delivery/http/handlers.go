package http

import (
	"electronic_diary/app/internal/domain/user"
	"electronic_diary/app/internal/domain/user/dto"
	"electronic_diary/app/internal/utils/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	userUC user.UseCase
}

func NewHandlers(userUC user.UseCase) user.Handlers {
	return userHandlers{userUC: userUC}
}

func (u userHandlers) Registration(c *gin.Context) {
	body, err := json.ValidateAndParseJSON[dto.UserCreateDTO](c)
	if err != nil {
		return
	}

	currentUser, err := u.userUC.Create(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"email": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, currentUser)
}

func (u userHandlers) Authorization(c *gin.Context) {
	body, err := json.ValidateAndParseJSON[dto.UserAuthorizationDTO](c)
	if err != nil {
		return
	}

	candidateUser, err := u.userUC.Authorization(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong login or password",
		})
		return
	}

	c.JSON(http.StatusOK, candidateUser)
}
