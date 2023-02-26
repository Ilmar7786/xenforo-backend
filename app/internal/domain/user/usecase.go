package user

import (
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
)

type UseCase interface {
	Registration(userDto dto.UserCreateDTO) (*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, bool)
	Authorization(userDto dto.UserAuthorizationDTO) (*model.User, error)
	BanUser(userDTO dto.UserBanDTO) (bool, error)
	ActivateEmail(linkID string) (bool, error)
}
