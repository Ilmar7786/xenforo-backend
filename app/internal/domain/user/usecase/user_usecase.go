package usecase

import (
	"errors"
	"fmt"

	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"

	"gorm.io/gorm"
)

type UserUC struct {
	db *gorm.DB
}

func NewUserUseCase(db *gorm.DB) user.UseCase {
	return &UserUC{db: db}
}

func (u *UserUC) Create(userDto dto.UserCreateDTO) (*model.User, error) {
	_, isUser := u.FindByEmail(userDto.Email)
	if !isUser {
		return nil, errors.New(fmt.Sprintf("user with mail %s already exists", userDto.Email))
	}

	hashedPassword, err := hashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	currentUser := &model.User{
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: hashedPassword,
	}
	u.db.Create(currentUser)

	return currentUser, nil
}

func (u *UserUC) FindAll() ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
	return nil, nil
}

func (u *UserUC) FindByID(id string) (*model.User, error) {
	var candidateUser model.User
	result := u.db.Where("id = ?", id).First(&candidateUser)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &candidateUser, nil
}

func (u *UserUC) Update(id string, userDto dto.UserUpdateDTO) (*model.User, error) {
	//TODO implement me
	panic("implement me")
	return nil, nil
}

func (u *UserUC) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
	return false, nil
}

func (u *UserUC) FindByEmail(email string) (*model.User, bool) {
	var currentUser model.User
	result := u.db.Where("email = ?", email).First(&currentUser)

	if result.RowsAffected == 0 {
		return nil, true
	}

	return &currentUser, false
}

func (u *UserUC) Authorization(userDto dto.UserAuthorizationDTO) (*model.User, error) {
	candidateUser, isUser := u.FindByEmail(userDto.Email)
	if isUser {
		return nil, errors.New(fmt.Sprintf("user with mail %s not found", userDto.Email))
	}

	isPasswordHash := checkPasswordHash(candidateUser.Password, userDto.Password)
	if !isPasswordHash {
		return nil, errors.New("incorrect password")
	}

	return candidateUser, nil
}

func (u *UserUC) BanUser(userID string) (bool, error) {
	//TODO implement me
	panic("implement me")
	return false, nil
}
