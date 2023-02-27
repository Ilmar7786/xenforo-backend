package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserRegistrationDTO struct {
	Email               string `json:"email" example:"example@mail.ru" minLength:"5" maxLength:"40"`
	Name                string `json:"name" example:"Иван" minLength:"2" maxLength:"20"`
	Password            string `json:"password" example:"12345678" minLength:"8" maxLength:"18"`
	RedirectActiveEmail string `json:"redirectActiveEmail" example:"https://example.ru/email/activate"`
}

func (u UserRegistrationDTO) Validate() error {
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
		validation.Field(&u.RedirectActiveEmail,
			validation.Required,
			is.URL,
		),
	)
}
