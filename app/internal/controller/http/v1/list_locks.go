package v1

import (
	"context"
	"net/http"
	"xenforo/app/internal/domain/auth/middleware"
	"xenforo/app/internal/domain/list_lock"
	"xenforo/app/internal/domain/list_lock/dto"
	"xenforo/app/pkg/api/validate"

	"github.com/gin-gonic/gin"
)

type listLockRoutes struct {
	ctx        context.Context
	listLockUC list_lock.UseCase
}

func newListLockRouters(handler *gin.RouterGroup, ctx context.Context, authMiddleware middleware.Init, listLockUC list_lock.UseCase) {
	r := listLockRoutes{
		ctx:        ctx,
		listLockUC: listLockUC,
	}

	h := handler.Group("/list-locks")
	{
		h.POST("/add", authMiddleware.AdminAuth(), r.Add)
		h.POST("/remove", authMiddleware.AdminAuth(), r.Remove)
		h.GET("/check-block", r.CheckBlock)
	}
}

func (l listLockRoutes) Add(c *gin.Context) {
	body, err := validate.ParseAndValidateJSON[dto.ListLockAddDTO](c)
	if err != nil {
		return
	}

	add, err := l.listLockUC.Add(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, add)
}

func (l listLockRoutes) Remove(c *gin.Context) {

}

func (l listLockRoutes) CheckBlock(c *gin.Context) {
	userIP := c.ClientIP()
	blocked, err := l.listLockUC.FindByIP(userIP)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blocked": blocked,
		"your_ip": userIP,
	})
}
