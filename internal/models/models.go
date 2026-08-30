package models

type TelegramRequest struct {
	Token   string
	ChatID  int64
	Message string
	IP      string
}
