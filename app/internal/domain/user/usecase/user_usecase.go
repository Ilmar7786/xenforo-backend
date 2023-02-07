package usecase

import (
	"electronic_diary/app/internal/domain/user"
	"electronic_diary/app/internal/domain/user/dto"
	"electronic_diary/app/internal/domain/user/model"

	"gorm.io/gorm"
)

type UserUC struct {
	db *gorm.DB
}

func NewUserUseCase(db *gorm.DB) user.UseCase {
	return &UserUC{db: db}
}

func (u *UserUC) Create(userDto dto.UserCreateDTO) (*model.User, error) {
	return nil, nil
}

func (u *UserUC) FindAll() ([]*model.User, error) {
	return nil, nil
}

func (u *UserUC) FindByID(id string) (*model.User, error) {
	return nil, nil
}

func (u *UserUC) Update(id string, userDto dto.UserUpdateDTO) (*model.User, error) {
	return nil, nil
}

func (u *UserUC) Delete(id string) error {
	return nil
}
