package handler

import (
	"log"
	"net/http"

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

func (s *Server) callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqests, err := s.Bot.ParseRequest(r)
	for _, req := range reqests {
		var response linebot.SendingMessage
		switch req.Type {
		case linebot.EventTypeMessage:
			response = s.getMessageRequestAction(ctx, req)
		case linebot.EventTypeFollow:
			response = s.getMessageFollowAction(ctx, req)
		}
		// 全処理をここで一括
		if _, err = s.Bot.ReplyMessage(req.ReplyToken, response).Do(); err != nil {
			log.Println(err)
		}
	}
}
