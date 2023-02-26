package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserCreateDTO struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	WhereSendLink string `json:"whereSendLink"`
}

func (u UserCreateDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(&u.Name,
			validation.Required,
			validation.Length(2, 20),
		),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 18),
		),
		validation.Field(&u.WhereSendLink,
			validation.Required,
			is.URL,
		),
	)
}
