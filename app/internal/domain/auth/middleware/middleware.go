package middleware

import (
	"context"

	"xenforo/app/internal/domain/user"
)

type Init struct {
	privateKey string
	ctx        context.Context
	userUC     user.UseCase
}

func NewAuth(ctx context.Context, privateKey string, userUC user.UseCase) Init {
	return Init{
		privateKey: privateKey,
		ctx:        ctx,
		userUC:     userUC,
	}
}
