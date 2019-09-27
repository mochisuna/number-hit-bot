package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mochisuna/number-hit-bot/application"
	"github.com/mochisuna/number-hit-bot/config"
	"github.com/mochisuna/number-hit-bot/handler"
	"github.com/mochisuna/number-hit-bot/infrastructure"
	"github.com/mochisuna/number-hit-bot/infrastructure/database"
)

func main() {
	// parse options
	path := flag.String("c", "_tools/local/config.toml", "config file")
	flag.Parse()

	// import config
	conf := &config.Config{}
	if err := config.New(conf, *path); err != nil {
		panic(err)
	}
	// init firebase client
	firestoreClient, err := database.NewFirestore(&conf.Firestore)
	if err != nil {
		panic(err)
	}
	defer firestoreClient.Close()
	// init repository
	userRepo := infrastructure.NewUserRepository(firestoreClient)
	// init service
	callbackService := application.NewCallbackService(userRepo)

	service := &handler.Services{
		CallbackService: callbackService,
	}

	// init line bot
	bot, err := handler.NewLineBot(&conf.LineBot)
	if err != nil {
		panic(err)
	}

	// 依存関係は全てここでinjection
	server := handler.New(conf.Server.Port, service, bot)
	log.Println("Start server")
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed ListenAndServe. err: %v", err))
	}
}
