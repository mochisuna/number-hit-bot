package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mochisuna/number-hit-bot/domain"
)

// botのアクションのみを統括
const (
	ActionEventStart  = "Start"
	ActionEventCancel = "Cancel"
)

func (s *Server) getMessageFollowAction(ctx context.Context, req *linebot.Event) linebot.SendingMessage {
	profile, err := s.Bot.GetProfile(req.Source.UserID).Do()
	if err != nil {
		log.Printf("error reason: %#v", err.Error())
		return linebot.NewTextMessage("プロフィール参照時にエラーが発生しました")
	}

	if err := s.CallbackService.Reset(ctx, domain.UserID(req.Source.UserID)); err != nil {
		log.Printf("error ins Start: %v", err.Error())
		return linebot.NewTextMessage("意図しない処理が発生しました。すまんの。")
	}
	return linebot.NewTextMessage(profile.DisplayName + "さん、\n登録ありがとう！\n早速だけどゲームスタート！\n僕が考えている数字が何かを当ててね！\n1から100のどれかだよ！")
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
func (s *Server) getMessageRequestAction(ctx context.Context, req *linebot.Event) linebot.SendingMessage {
	switch message := req.Message.(type) {
	case *linebot.TextMessage:
		switch message.Text {
		case ActionEventStart:
			if err := s.CallbackService.Reset(ctx, domain.UserID(req.Source.UserID)); err != nil {
				log.Printf("error in ActionEventStart: %v", err.Error())
				return linebot.NewTextMessage("意図しない処理が発生しました。すまんの。")
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
