package usecase

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/sport"
	"xenforo/app/internal/domain/sport/model"
	"xenforo/app/pkg/client/flashliveSports"
	"xenforo/app/pkg/logging"
)

const baseRoute = "/sports"

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

func (s *SportsUC) NumberSportEvents(sportID string) (*model.SportData, error) {
	var sports model.SportData

	s.flashliveSports.AddQuery("sport_id", sportID)
	err := s.flashliveSports.Request(http.MethodGet, baseRoute+"/events-count", &sports)
	if err != nil {
		logging.Error(s.ctx, err)
		return nil, err
	}

	return &sports, nil
}
