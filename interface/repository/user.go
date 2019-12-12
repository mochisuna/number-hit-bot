package infrastructure

import (
	"context"
	"log"
	"math/rand"

	"github.com/mochisuna/number-hit-bot/domain"
	"github.com/mochisuna/number-hit-bot/domain/repository"
	"github.com/mochisuna/number-hit-bot/infrastructure/firebase"
)

type userRepository struct {
	client *firebase.FirestoreClient
}

func NewUserRepository(client *firebase.FirestoreClient) repository.UserRepository {
	return &userRepository{
		client,
	}
}
func (r *userRepository) GetUser(ctx context.Context, id domain.UserID) (*domain.User, error) {
	dsnap, err := r.client.Collection("users").Doc(string(id)).Get(ctx)
	if err != nil {
		log.Printf("error in GetUser: %v\n", err.Error())
		return nil, err
	}
	log.Printf("%#v\n", dsnap.Data())
	res := domain.User{}

	if err = dsnap.DataTo(&res); err != nil {
		log.Printf("error in GetUser: %v\n", err.Error())
	}
	return &res, nil
}

func (r *userRepository) ResetUser(ctx context.Context, id domain.UserID) (*domain.User, error) {
	num := rand.Intn(100)
	user := domain.User{
		ID:        id,
		Answer:    domain.AnswerNumber(num + 1),
		MissCount: 0,
	}
	_, err := r.client.Collection("users").Doc(string(user.ID)).Set(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	_, err := r.client.Collection("users").Doc(string(user.ID)).Set(ctx, user)
	return err
}

