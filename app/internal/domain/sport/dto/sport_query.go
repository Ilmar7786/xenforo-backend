package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SportQueryDTO struct {
	SportID string `json:"sportId"`
}

func (s SportQueryDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.SportID, validation.Required),
	)
}
