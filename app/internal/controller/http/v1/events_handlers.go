package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/event"
	"xenforo/app/internal/domain/event/dto"
	"xenforo/app/pkg/client/flashliveSports"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	ctx     context.Context
	eventUC event.UseCase
}

func newEventHandler(ctx context.Context, eventUC event.UseCase) *eventHandler {
	return &eventHandler{
		ctx:     ctx,
		eventUC: eventUC,
	}
}

// GetList godoc
// @Summary Список живых событий
// @Tags 	events
// @Accept 	json
// @Produce json
// @Param 	sportID		 query string false "id спорта"
// @Param 	timeZone 	 query string false "Часовой пояс"
// @Param 	locale 		 query string false "ID спорта"
// @Success 200 {object} model.LiveEventData
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /events [get]
func (s *eventHandler) GetList(c *gin.Context) {
	var input dto.LiveEventDTO
	input.Locale = c.Query(flashliveSports.QueryParamLocale)
	input.TimeZone = c.Query(flashliveSports.QueryTimeParamZone)
	input.SportId = c.Query(flashliveSports.QueryParamSport)

	res, err := s.eventUC.GetLiveEventsList(flashliveSports.Filter{
		Locale:   input.Locale,
		TimeZone: input.TimeZone,
		SportID:  input.SportId,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
