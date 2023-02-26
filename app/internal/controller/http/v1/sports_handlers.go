package v1

import (
	"context"
	"net/http"

	"xenforo/app/internal/domain/sport"

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

func (s *sportHandler) GetList(c *gin.Context) {
	res := s.sportUC.NumberSportEvents()
	c.JSON(http.StatusOK, res)
}
