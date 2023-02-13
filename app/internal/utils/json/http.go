package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateAndParseJSON[T validation.Validatable](c *gin.Context) (T, error) {
	var body T
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return body, err
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return body, err
	}

	return body, nil
}
