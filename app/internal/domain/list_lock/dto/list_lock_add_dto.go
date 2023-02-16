package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ListLockAddDTO struct {
	IP     string `json:"ip"`
	UserID string `json:"userID"`
}

func (l ListLockAddDTO) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.IP, validation.Required, is.IP),
		validation.Field(&l.UserID),
	)
}
