package application

import (
	"context"
	"log"

	"github.com/mochisuna/number-hit-bot/domain"
	"github.com/mochisuna/number-hit-bot/domain/repository"
	"github.com/mochisuna/number-hit-bot/domain/service"
)

type CallbackService struct {
	userRepo repository.UserRepository
}

// NewCallbackService inject eventRepo
func NewCallbackService(userRepo repository.UserRepository) service.CallbackService {
	return &CallbackService{
		userRepo: userRepo,
	}
}

func (s *CallbackService) Follow(ctx context.Context, id domain.UserID) error {
	user, err := s.userRepo.ResetUser(ctx, id)
	log.Println(user)
	return err
}
func (s *CallbackService) Reset(ctx context.Context, id domain.UserID) error {
	user, err := s.userRepo.ResetUser(ctx, id)
	log.Println(user)
	return err
}
func (s *CallbackService) Check(ctx context.Context, id domain.UserID, ans domain.AnswerNumber) (domain.EventStatus, error) {
	user, err := s.userRepo.GetUser(ctx, id)
	if err != nil {
		return -1, err
	}
	log.Println(user)
	if user.Answer == -1 {
		return domain.NODATA, nil
	}
	if user.Answer == ans {
		user.Answer = -1
		user.MissCount = 0
		err = s.userRepo.UpdateUser(ctx, user)
		return domain.CLEAR, err
	}
	user.MissCount++
	if user.MissCount > domain.MAXIMUM_MISSCOUNT {
		user.Answer = -1
		user.MissCount = 0
		err = s.userRepo.UpdateUser(ctx, user)
		return domain.GAMEOVER, err
	}
	if user.Answer < ans {
		err = s.userRepo.UpdateUser(ctx, user)
		return domain.FAIL_TOO_LARGE, nil
	}
	err = s.userRepo.UpdateUser(ctx, user)
	return domain.FAIL_TOO_SMALL, nil
}
