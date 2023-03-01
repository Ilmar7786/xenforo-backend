package usecase

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/event"
	"xenforo/app/internal/domain/event/model"
	"xenforo/app/pkg/client/flashliveSports"
)

const baseRoute = "/events"

type EventUC struct {
	ctx             context.Context
	flashliveSports *flashliveSports.Client
}

func NewEventUseCase(ctx context.Context, client *flashliveSports.Client) event.UseCase {
	return &EventUC{
		ctx:             ctx,
		flashliveSports: client,
	}
}

func (e EventUC) GetLiveEventsList(filter flashliveSports.Filter) (*model.LiveEventData, error) {
	var events model.LiveEventData

	e.flashliveSports.AddQuery(flashliveSports.QueryParamLocale, filter.Locale)
	e.flashliveSports.AddQuery(flashliveSports.QueryTimeParamZone, filter.TimeZone)
	e.flashliveSports.AddQuery(flashliveSports.QueryParamSport, filter.SportID)
	err := e.flashliveSports.Request(http.MethodGet, baseRoute+"/live-list", &events)
	if err != nil {
		return nil, err
	}

	return &events, nil
}
