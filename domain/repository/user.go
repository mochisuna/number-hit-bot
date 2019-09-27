package repository

import (
	"context"

	"github.com/mochisuna/number-hit-bot/domain"
)

type UserRepository interface {
	ResetUser(context.Context, domain.UserID) (*domain.User, error)
	UpdateUser(context.Context, *domain.User) error
	GetUser(context.Context, domain.UserID) (*domain.User, error)
}
