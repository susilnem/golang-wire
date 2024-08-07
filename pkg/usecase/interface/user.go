package interfaces

import (
	"context"
	"go-wire/pkg/domain"
)

type UserUseCase interface {
	SignUpUser(ctx context.Context, user domain.User) (domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)

	// generate token
	GenerateAccessToken(ctx context.Context, tokenParams GenerateTokenParams) (tokenString string, err error)
	GenerateRefreshToken(ctx context.Context, tokenParams GenerateTokenParams) (tokenString string, err error)
}

type GenerateTokenParams struct {
	UserID   uint
}
