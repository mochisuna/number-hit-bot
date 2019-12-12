package interactor

import (
	"context"
	"log"

	"github.com/mochisuna/number-hit-bot/"
	"github.com/mochisuna/number-hit-bot/usecase/repository"
	"github.com/mochisuna/number-hit-bot/usecase/presenter"
	"github.com/mochisuna/number-hit-bot/interface/interactor"
)

type linebotInteractor struct {
	userRepo repository.UserRepository
	botRepo repository.BotRepository
	botPres presenter.BotPresenter
}

// NewLineBotInteractor inject eventRepo
func NewLineBotInteractor(userRepo repository.UserRepository) interactor.linebotInteractor {
	return &linebotInteractor{
		userRepo: userRepo,
	}
}
func (s *linebotInteractor)Reply(r *http.Request) {
	ctx := r.Context()
	reqests, err := c.botRepo.ParseRequest(r)
	for _, req := range reqests {
		switch req.Type {
		case linebot.EventTypeMessage:
			request(ctx, req)
		case linebot.EventTypeFollow:
			follow(ctx, req)
		}
		s.botPres.Response()
	}
}

func (s *Server) follow(ctx context.Context, req *linebot.Event) string {
	profile, err := s.botRepo.GetUserName(req.Source.UserID)
	if err := s.userRepo.Reset(ctx, domain.UserID(req.Source.UserID)); err != nil {
		log.Printf("error ins Start: %v", err.Error())
		return "意図しない処理が発生しました。すまんの。"
	}
	return profile + "さん、\n登録ありがとう！\n早速だけどゲームスタート！\n僕が考えている数字が何かを当ててね！\n1から100のどれかだよ！"
}


func getNum(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	if num < 1 || num > 100 {
		return -1, nil
	}
	return num, nil
}

func request(ctx context.Context, req *linebot.Event) {
	switch message := req.Message.(type) {
	case *linebot.TextMessage:
		switch message.Text {
		case ActionEventStart:
			if err := s.CallbackService.Reset(ctx, domain.UserID(req.Source.UserID)); err != nil {
				log.Printf("error in ActionEventStart: %v", err.Error())
				return "意図しない処理が発生しました。すまんの。"
			}
			return linebot.NewTextMessage("僕が考えている数字が何かを当ててね！\n1から100のどれかだよ！")
		case ActionEventCancel:
			return linebot.NewTextMessage("残念・・・")
		default:
			ask := "ゲームをはじめる？"
			num, err := getNum(message.Text)
			if err != nil {
				log.Printf("error in getNum: %v", err.Error())
				return linebot.NewTextMessage("あれ？ 1〜100の数字を入力してね！")
			} else if num == -1 {
				return linebot.NewTextMessage("1〜100の範囲だよ！")
			}
			status, err := s.CallbackService.Check(ctx, domain.UserID(req.Source.UserID), domain.AnswerNumber(num))
			if err != nil {
				return linebot.NewTextMessage("意図しない処理が発生しました。すまんの。")
			}
			switch status {
			case domain.NODATA:
				ask = "まだ数字を考えてなかったよ。\nそれじゃあ" + ask
			case domain.CLEAR:
				ask = fmt.Sprintf("すごい！正解は%vでした！\nもう一度%v", num, ask)
			case domain.FAIL_TOO_LARGE:
				return linebot.NewTextMessage("残念！もっと小さい数字だよ！")
			case domain.FAIL_TOO_SMALL:
				return linebot.NewTextMessage("残念！もっと大きい数字だよ！")
			case domain.GAMEOVER:
				ask = fmt.Sprintf("残念！%v回以上間違ったからゲームオーバーだ！\nもう一度%v", domain.MAXIMUM_MISSCOUNT, ask)
			}
			return linebot.NewTemplateMessage(
				"start event",
				linebot.NewConfirmTemplate(
					ask,
					linebot.NewMessageAction("やろう！", ActionEventStart),
					linebot.NewMessageAction("やめとく", ActionEventCancel),
				),
			)
		}
	}
	return linebot.NewTextMessage("意図しない処理が発生しました。すまんの。")
}

// func (s *linebotInteractor) Follow(ctx context.Context, id domain.UserID) error {
// 	user, err := s.userRepo.ResetUser(ctx, id)
// 	log.Println(user)
// 	return err
// }
// func (s *linebotInteractor) Reset(ctx context.Context, id domain.UserID) error {
// 	user, err := s.userRepo.ResetUser(ctx, id)
// 	log.Println(user)
// 	return err
// }
// func (s *linebotInteractor) Check(ctx context.Context, id domain.UserID, ans domain.AnswerNumber) (domain.EventStatus, error) {
// 	user, err := s.userRepo.GetUser(ctx, id)
// 	if err != nil {
// 		return -1, err
// 	}
// 	log.Println(user)
// 	if user.Answer == -1 {
// 		return domain.NODATA, nil
// 	}
// 	if user.Answer == ans {
// 		user.Answer = -1
// 		user.MissCount = 0
// 		err = s.userRepo.UpdateUser(ctx, user)
// 		return domain.CLEAR, err
// 	}
// 	user.MissCount++
// 	if user.MissCount > domain.MAXIMUM_MISSCOUNT {
// 		user.Answer = -1
// 		user.MissCount = 0
// 		err = s.userRepo.UpdateUser(ctx, user)
// 		return domain.GAMEOVER, err
// 	}
// 	if user.Answer < ans {
// 		err = s.userRepo.UpdateUser(ctx, user)
// 		return domain.FAIL_TOO_LARGE, nil
// 	}
// 	err = s.userRepo.UpdateUser(ctx, user)
// 	return domain.FAIL_TOO_SMALL, nil
// }

// func (s *linebotInteractor)ParseRequest(context.Context) error{
// 	return nil
// }

// func (s *linebotInteractor)GetRequestMessage(context.Context) error{
// 	return nil
// }

// func (s *linebotInteractor)GetFollowMessage(context.Context) error{
// 	return nil
// }


