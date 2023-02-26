package usecase

import (
	"context"
	"fmt"

	"xenforo/app/internal/config"
	"xenforo/app/internal/domain/mail"
	"xenforo/app/internal/domain/mail/model"
	"xenforo/app/pkg/logging"
	"xenforo/app/pkg/mailer"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MailUC struct {
	ctx    context.Context
	db     *gorm.DB
	mailer *mailer.Mailer
}

func NewMailUseCase(ctx context.Context, cfg *config.Config, db *gorm.DB) mail.UseCase {
	mailerConfig := mailer.Config{
		From:     cfg.Mail.From,
		HOST:     cfg.Mail.Host,
		Port:     cfg.Mail.Port,
		Username: cfg.Mail.Username,
		Password: cfg.Mail.Password,
		SSL:      cfg.Mail.SSL,
	}

	newMailer, err := mailer.NewMailer(mailerConfig)
	if err != nil {
		logging.Error(ctx, err)
	}

	return &MailUC{
		ctx:    ctx,
		mailer: newMailer,
		db:     db,
	}
}

func (m MailUC) Activate(linkID string) (bool, error) {
	mailActivate, err := m.FindById(linkID)
	if err != nil {
		return false, err
	}

	if mailActivate.ID == "" {
		return false, nil
	}

	resultDelete, err := m.Delete(mailActivate.ID)
	if err != nil {
		return false, err
	} else if !resultDelete {
		return false, nil
	}

	return true, nil
}

func (m MailUC) GenerateActivateLink(userID, to, link string) error {
	m.db.Where("user_id = ?", userID).Delete(model.MailActivate{})

	linkUUID := uuid.New().String()
	linkURL := fmt.Sprintf("%s/%s", link, linkUUID)
	message := "Перейдите по ссылки для подтверждения email " + linkURL

	err := m.mailer.NewMessage(to, "Регистрация аккаунта", message)
	if err != nil {
		return err
	}

	mailMessage := model.MailActivate{
		ID:     linkUUID,
		UserID: userID,
	}
	m.db.Create(&mailMessage)

	return nil
}

func (m MailUC) FindById(linkID string) (*model.MailActivate, error) {
	var currentMail model.MailActivate

	result := m.db.Where("id = ?", linkID).First(&currentMail)
	if result.Error != nil {
		return nil, result.Error
	}

	return &currentMail, nil
}

func (m MailUC) FindByUserID(userID string) (*model.MailActivate, error) {
	var currentMail model.MailActivate
	result := m.db.Where("user_id = ?", userID).First(&currentMail)

	if result.Error != nil {
		return nil, result.Error
	}

	return &currentMail, nil
}

func (m MailUC) Delete(linkID string) (bool, error) {
	result := m.db.Where("id = ?", linkID).Delete(&model.MailActivate{})

	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
