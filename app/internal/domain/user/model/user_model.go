package model

import (
	"time"
	"xenforo/app/internal/domain/list_lock/model"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Name      string    `json:"name" gorm:"not null"`
	IsAdmin   bool      `json:"isAdmin" gorm:"default: false; not null"`
	IsBanned  bool      `json:"isBanned"  gorm:"default: false; not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`

	ListLocks []model.ListLock `json:"-" gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "users"
}
