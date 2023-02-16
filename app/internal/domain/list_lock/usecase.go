package list_lock

import (
	"xenforo/app/internal/domain/list_lock/dto"
	"xenforo/app/internal/domain/list_lock/model"
)

type UseCase interface {
	Add(dto dto.ListLockAddDTO) (bool, error)
	Remove(IP, userID string) (bool, error)
	GetAll() []model.ListLock
	FindByIP(ip string) (bool, error)
}
