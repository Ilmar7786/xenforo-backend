package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserUpdateDTO struct {
	Email    string `json:"email" minLength:"5" maxLength:"40"`
	Name     string `json:"name" minLength:"2" maxLength:"20"`
	Password string `json:"password" minLength:"5" maxLength:"18"`
}

func (u UserUpdateDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required,
			is.Email,
			validation.Length(5, 40),
		),
		validation.Field(&u.Name,
			validation.Required,
			validation.Length(2, 20),
		),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 18),
		),
	)
}
