package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserAuthorizationDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserAuthorizationDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 18),
		),
	)
}
