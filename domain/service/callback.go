package service

import (
	"context"

	"github.com/mochisuna/number-hit-bot/domain"
)

type CallbackService interface {
	Follow(context.Context, domain.UserID) error
	Reset(context.Context, domain.UserID) error
	Check(context.Context, domain.UserID, domain.AnswerNumber) (domain.EventStatus, error)
}
