package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mochisuna/number-hit-bot/application"
	"github.com/mochisuna/number-hit-bot/config"
	"github.com/mochisuna/number-hit-bot/handler"
	"github.com/mochisuna/number-hit-bot/infrastructure"
	"github.com/mochisuna/number-hit-bot/infrastructure/firebase"
)

func main() {
	// parse options
	env := flag.String("e", "local", "environment")
	flag.Parse()

	// import config
	conf, err := config.New(*env)
	if err != nil {
		panic(fmt.Sprintf("Loading config failed. err: %+v", err))
	}

	// init firebase client
	firestoreClient, err := firebase.NewFirestore(&conf.Firestore)
	if err != nil {
		panic(fmt.Sprintf("Loading firestore-cilent failed. err: %+v", err))
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
