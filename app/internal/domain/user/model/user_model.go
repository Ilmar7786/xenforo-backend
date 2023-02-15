package model

import (
	"time"
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
}

func (u *User) TableName() string {
	return "users"
}
