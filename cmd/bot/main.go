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

	firestoreClient, err := database.NewFirestore(&conf.Firestore)
	if err != nil {
		panic(err)
	}
	defer firestoreClient.Close()

	userRepo := infrastructure.NewUserRepository(firestoreClient)
	callbackService := application.NewCallbackService(userRepo)
	service := &handler.Services{
		CallbackService: callbackService,
	}

	bot, err := handler.NewLineBot(&conf.LineBot)
	if err != nil {
		panic(err)
	}
	server := handler.New(conf.Server.Port, service, bot)
	log.Println("Start server")
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed ListenAndServe. err: %v", err))
	}
}
