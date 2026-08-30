package tServ

import (
	"Service/internal/models"

	tele "gopkg.in/telebot.v4"
)

type TGService struct {
}

func New() *TGService {
	return &TGService{}
}

func (tb *TGService) Send(req models.WebhookReceiver) error {
	bot, err := tele.NewBot(tele.Settings{
		Token: req.Token,
	})

	if err != nil {
		return err
	}

	_, err = bot.Send(&tele.Chat{ID: req.ChatID}, req.Message)

	return err
}
