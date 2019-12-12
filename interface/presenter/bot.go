package presenter

import (
	"context"
)

type BotPresenter interface {
	ReplyMessage(context.Context)
}
