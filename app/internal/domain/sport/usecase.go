package sport

import "xenforo/app/internal/domain/sport/model"

type UseCase interface {
	NumberSportEvents(sportID string) (*model.SportData, error)
}
