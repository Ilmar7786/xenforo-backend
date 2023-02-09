package dto

import validation "github.com/go-ozzo/ozzo-validation"

type UserUpdateDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u UserUpdateDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required.Error("Введите почту"),
		),
		validation.Field(&u.Name,
			validation.Required.Error("Введите имя"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("Введите Пароль"),
			validation.Length(8, 18).Error("Пароль должен содержать от 8 до 18 символов"),
		),
	)
}
