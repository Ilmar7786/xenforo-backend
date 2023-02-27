package user

import (
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
)

type UseCase interface {
	Registration(userDto dto.UserRegistrationDTO) (*model.UserAndTokens, error)
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) *model.User
	Authorization(userDto dto.UserAuthorizationDTO) (*model.UserAndTokens, error)
	BanUser(userDTO dto.UserBanDTO) (bool, error)
	ActivateEmail(linkID string) (bool, error)
	FindAll() []*model.User
}
