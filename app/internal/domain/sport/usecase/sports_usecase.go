package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"xenforo/app/internal/domain/sport"
	"xenforo/app/internal/domain/sport/model"
	"xenforo/app/pkg/client/flashliveSports"
	"xenforo/app/pkg/logging"
)

type (
	SportsUC struct {
		ctx             context.Context
		flashliveSports *flashliveSports.Client
	}
)

func NewSportsUseCase(ctx context.Context, client *flashliveSports.Client) sport.UseCase {
	return &SportsUC{
		ctx:             ctx,
		flashliveSports: client,
	}
}

func (s *SportsUC) NumberSportEvents() *model.SportData {
	res, err := s.flashliveSports.Request(http.MethodGet, "sports/events-count", &flashliveSports.Filter{
		Locale:   "ru_RU",
		TimeZone: "+3",
	})
	if err != nil {
		logging.Error(s.ctx, err)
		return nil
	}

	var sports model.SportData
	err = json.Unmarshal(res, &sports)
	if err != nil {
		logging.Error(s.ctx, err)
		return nil
	}

	return &sports
}
