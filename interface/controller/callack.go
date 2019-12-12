package controller

import (
	"net/http"

	"github.com/mochisuna/number-hit-bot/interface/interactor"
)

const ctrkey = "admin"

type callbackController struct {
	linebotInteractor interactor.LineBotInteractor
}

func NewCallbackController(linebotInteractor interactor.LineBotInteractor) controller.CallbackController {
	return &callbackController{
		linebotInteractor,
	}
}

func (c *callbackController) Callback(w http.ResponseWriter, r *http.Request) {
	c.linebotInteractor.Reply(r)
}
