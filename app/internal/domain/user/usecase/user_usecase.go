package usecase

import (
	"context"
	"errors"
	"fmt"

	"xenforo/app/internal/domain/mail"
	"xenforo/app/internal/domain/user"
	"xenforo/app/internal/domain/user/dto"
	"xenforo/app/internal/domain/user/model"
	"xenforo/app/pkg/logging"

	"gorm.io/gorm"
)

type UserUC struct {
	db     *gorm.DB
	mailUC mail.UseCase
	ctx    context.Context
}

func NewUserUseCase(ctx context.Context, db *gorm.DB, mailUC mail.UseCase) user.UseCase {
	return &UserUC{
		db:     db,
		mailUC: mailUC,
		ctx:    ctx,
	}
}

func (u *UserUC) Registration(userDto dto.UserCreateDTO) (*model.User, error) {
	_, isUser := u.FindByEmail(userDto.Email)
	if !isUser {
		return nil, errors.New(fmt.Sprintf("user with mail %s already exists", userDto.Email))
	}

	hashedPassword, err := hashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	currentUser := model.User{
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: hashedPassword,
	}
	result := u.db.Create(&currentUser)
	if result.Error != nil {
		return nil, result.Error
	}

	err = u.mailUC.GenerateActivateLink(currentUser.ID, currentUser.Email, userDto.WhereSendLink)
	if err != nil {
		logging.Error(u.ctx, err)
	}

	return &currentUser, nil
}

func (u *UserUC) FindByID(id string) (*model.User, error) {
	var candidateUser model.User
	result := u.db.Where("id = ?", id).First(&candidateUser)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &candidateUser, nil
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

func (u *UserUC) BanUser(userDto dto.UserBanDTO) (bool, error) {
	currentUser, err := u.FindByID(userDto.UserID)
	if err != nil {
		return false, err
	}
	if currentUser.ID == "" {
		return false, nil
	}

	currentUser.IsBanned = userDto.IsBan
	result := u.db.Save(&currentUser)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (u *UserUC) ActivateEmail(linkID string) (bool, error) {
	activate, err := u.mailUC.Activate(linkID)
	if err != nil {
		return false, err
	}
	if activate {
		return false, nil
	}

	return true, nil
}
