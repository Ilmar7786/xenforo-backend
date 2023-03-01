package event

import (
	"xenforo/app/internal/domain/event/model"
	"xenforo/app/pkg/client/flashliveSports"
)

type UseCase interface {
	GetLiveEventsList(filter flashliveSports.Filter) (*model.LiveEventData, error)
}
