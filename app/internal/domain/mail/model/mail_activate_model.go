package model

import (
	"time"
)

type MailActivate struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m MailActivate) TableName() string {
	return "mail_activates"
}
