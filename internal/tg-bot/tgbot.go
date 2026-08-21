package tServ

import (
	tele "gopkg.in/telebot.v4"
)

type TGService struct {
}

func New() *TGService {
	return &TGService{}
}

func (tb *TGService) Send(token string, chatID int64, message string) error {
	bot, err := tele.NewBot(tele.Settings{
		Token: token,
	})

	if err != nil {
		return err
	}

	_, err = bot.Send(&tele.Chat{ID: chatID}, message)

	return err
}
