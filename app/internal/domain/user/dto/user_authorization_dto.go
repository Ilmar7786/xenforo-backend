package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserAuthorizationDTO struct {
	Email    string `json:"email" example:"example@mail.ru" minLength:"5" maxLength:"40"`
	Password string `json:"password" example:"12345678" minLength:"5" maxLength:"18"`
}

func (u UserAuthorizationDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required,
			is.Email,
			validation.Length(5, 40),
		),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 18),
		),
	)
}
