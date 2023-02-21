package validate

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ParseAndValidateJSON[T validation.Validatable](c *gin.Context) (T, error) {
	var body T
	if err := c.ShouldBindJSON(&body); err != nil {
		return body, err
	}

	if err := body.Validate(); err != nil {
		return body, err
	}

	return body, nil
}
