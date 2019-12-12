package linebot

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mochisuna/number-hit-bot/config"
)

// LineBot client
type LineBot struct {
	Bot *linebot.Client
}

// NewLineBot initialize client
func NewLineBot(config *config.LineBot) (*LineBot, error) {
	client, err := linebot.New(config.ChannelSecret, config.ChannelToken)
	if err != nil {
		return nil, err
	}

	return &LineBot{client}, nil
}
