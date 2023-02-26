package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserBanDTO struct {
	IsBan  bool   `json:"isBan"`
	UserID string `json:"-"`
}

func (u UserBanDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserID,
			validation.Required,
		),
	)
}
