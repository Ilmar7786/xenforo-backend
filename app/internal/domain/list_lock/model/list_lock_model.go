package model

import (
	"time"
)

type ListLock struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	IP        string    `gorm:"column:ip;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
	UserID    string    `gorm:"column:user_id; default:null"`
}

func (u *ListLock) TableName() string {
	return "list_locks"
}
