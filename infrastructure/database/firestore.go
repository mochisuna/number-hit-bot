package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/mochisuna/number-hit-bot/config"
	"google.golang.org/api/option"
)

type FirestoreClient struct {
	*firestore.Client
}

func NewFirestore(config *config.Firestore) (*FirestoreClient, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.JSONPath)
	conf := &firebase.Config{ProjectID: config.ProjectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Printf("error in NewApp")
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &FirestoreClient{client}, nil
}

func (c *FirestoreClient) Close() {
	c.Close()
}
