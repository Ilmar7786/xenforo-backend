package api

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ParseAndValidateJSON[T validation.Validatable](c *gin.Context) (T, error) {
	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, err
	}

	if err := input.Validate(); err != nil {
		return input, err
	}

	return input, nil
}
