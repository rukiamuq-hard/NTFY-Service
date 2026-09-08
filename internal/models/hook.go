package models

type RequestTelegram struct {
	Token   string `json:"token"`
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}

type NotificationLog struct {
	Recipient string `json:"recipient"`
	Provider  string `json:"provider"`
	Message   string `json:"message"`
	IP        string
}
