package repository

import (
	"context"

	"github.com/mochisuna/number-hit-bot/domain"
)

type UserRepository interface {
	GetUser(context.Context, domain.UserID) (*domain.User, error)
	ResetUser(context.Context, domain.UserID) (*domain.User, error)
	UpdateUser(context.Context, *domain.User) error
}
