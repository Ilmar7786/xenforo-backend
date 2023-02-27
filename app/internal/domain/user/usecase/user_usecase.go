package usecase

import (
	"context"
	"errors"
	"fmt"
	"xenforo/app/internal/config"
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
	cfg    *config.Config
}

func NewUserUseCase(ctx context.Context, cfg *config.Config, db *gorm.DB, mailUC mail.UseCase) user.UseCase {
	return &UserUC{
		db:     db,
		mailUC: mailUC,
		ctx:    ctx,
		cfg:    cfg,
	}
}

func (u *UserUC) Registration(userDto dto.UserRegistrationDTO) (*model.UserAndTokens, error) {
	existsUser := u.FindByEmail(userDto.Email)
	if existsUser != nil {
		return nil, errors.New(fmt.Sprintf("user with mail %s already existsUser", existsUser.Email))
	}

	hashedPassword, err := hashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: hashedPassword,
	}
	result := u.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	currentUser := model.UserAndTokens{
		User: newUser,
	}

	err = generateTokens(&currentUser, u.cfg)
	if err != nil {
		return nil, err
	}

	err = u.mailUC.GenerateActivateLink(currentUser.ID, currentUser.Email, userDto.RedirectActiveEmail)
	if err != nil {
		logging.Error(u.ctx, err)
		return nil, err
	}

	return &currentUser, nil
}

func (u *UserUC) FindByID(id string) (*model.User, error) {
	var currentUser model.User
	result := u.db.Where("id = ?", id).First(&currentUser)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &currentUser, nil
}

func (u *UserUC) FindByEmail(email string) *model.User {
	var currentUser model.User
	result := u.db.Where("email = ?", email).First(&currentUser)
	if result.RowsAffected == 0 {
		return nil
	}

	return &currentUser
}

func (u *UserUC) Authorization(userDto dto.UserAuthorizationDTO) (*model.UserAndTokens, error) {
	existsUser := u.FindByEmail(userDto.Email)
	if existsUser == nil {
		return nil, errors.New(fmt.Sprintf("user with mail %s not found", userDto.Email))
	}

	currentUser := &model.UserAndTokens{
		User: *existsUser,
	}

	isPasswordHash := checkPasswordHash(existsUser.Password, userDto.Password)
	if !isPasswordHash {
		return nil, errors.New("incorrect password")
	}

	err := generateTokens(currentUser, u.cfg)
	if err != nil {
		return nil, err
	}

	return currentUser, nil
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

func (u *UserUC) FindAll() []*model.User {
	users := make([]*model.User, 0)
	u.db.Find(&users)

	return users
}
