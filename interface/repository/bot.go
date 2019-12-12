package infrastructure

import (
	"context"
	"log"
	"math/rand"

	"github.com/mochisuna/number-hit-bot/domain"
	"github.com/mochisuna/number-hit-bot/usecase/repository"
	"github.com/mochisuna/number-hit-bot/infrastructure/linebot"
)

type botRepository struct {
	bot *linebot.LineBot
}

func NewBotRepository(bot *linebot.LineBot) repository.BotRepository {
	return &botRepository{
		bot,
	}
}
