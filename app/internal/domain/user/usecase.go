package user

import (
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
)

type UseCase interface {
	Create(userDto dto.UserCreateDTO) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindByID(id string) (*model.User, error)
	Update(id string, userDto dto.UserUpdateDTO) (*model.User, error)
	Delete(id string) (bool, error)
	FindByEmail(email string) (*model.User, bool)
	Authorization(userDto dto.UserAuthorizationDTO) (*model.User, error)
	BanUser(userID string) (bool, error)
}
