package usecase

import (
	"xenforo/app/internal/domain/list_lock"
	"xenforo/app/internal/domain/list_lock/dto"
	ListLockModel "xenforo/app/internal/domain/list_lock/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListLockUC struct {
	db *gorm.DB
}

func NewListLockUseCase(db *gorm.DB) list_lock.UseCase {
	return &ListLockUC{db: db}
}
func (l *ListLockUC) Add(dto dto.ListLockAddDTO) (bool, error) {
	listLock := ListLockModel.ListLock{
		IP:     dto.IP,
		UserID: dto.UserID,
	}

	err := l.db.Create(&listLock).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (l *ListLockUC) Remove(IP, userID string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *ListLockUC) GetAll() []ListLockModel.ListLock {
	var listLocks []ListLockModel.ListLock
	l.db.Preload(clause.Associations).Find(&listLocks)
	return listLocks
}

func (l *ListLockUC) FindByIP(ip string) (bool, error) {
	var listLock ListLockModel.ListLock
	result := l.db.Where("ip = ?", ip).First(&listLock)

	if result.RowsAffected == 0 {
		return false, nil
	}

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
