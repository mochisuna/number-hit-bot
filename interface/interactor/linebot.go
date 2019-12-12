package interactor

import (
	"net/http"
)

type LineBotInteractor interface {
	Reply(r *http.Request)
}

// Follow(context.Context, domain.UserID) error
// Reset(context.Context, domain.UserID) error
// Check(context.Context, domain.UserID, domain.AnswerNumber) (domain.EventStatus, error)
