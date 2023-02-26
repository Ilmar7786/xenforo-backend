package mail

import "xenforo/app/internal/domain/mail/model"

type UseCase interface {
	GenerateActivateLink(userID, to, link string) error
	Activate(linkID string) (bool, error)
	FindById(linkID string) (*model.MailActivate, error)
	FindByUserID(userID string) (*model.MailActivate, error)
	Delete(linkID string) (bool, error)
}
