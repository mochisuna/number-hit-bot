package repository

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type BotRepository interface {
	ParseRequest(*http.Request) ([]*linebot.Event, error)
	GetUserName(string) (string, error)
}
