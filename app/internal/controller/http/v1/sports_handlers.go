package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/sport"
	"xenforo/app/internal/domain/sport/dto"
	"xenforo/app/pkg/api"

	"github.com/gin-gonic/gin"
)

type sportHandler struct {
	ctx     context.Context
	sportUC sport.UseCase
}

func newSportHandler(ctx context.Context, sportUC sport.UseCase) *sportHandler {
	return &sportHandler{
		ctx:     ctx,
		sportUC: sportUC,
	}
}

// GetList godoc
// @Summary Список спортивных мероприятий
// @Tags sports
// @Description Получить данные о виде спорта и количестве спортивных событий
// @Accept 	json
// @Produce json
// @Param 	locale		query string false "Язык"
// @Success 200 {object} model.SportData
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /sports [get]
func (s *sportHandler) GetList(c *gin.Context) {
	input, err := api.ParseAndValidateJSON[dto.SportQueryDTO](c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := s.sportUC.NumberSportEvents(input.SportID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
