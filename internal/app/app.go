package app

import (
	"Service/internal/echo-serv"
)

type App struct {
	serv *echoserv.Server
}

func New() *App {
	return &App{}
}

func (a *App) Start() error {
	a.serv = echoserv.New()
	if err := a.serv.Start(":8080"); err != nil {
		return err
	}

	return nil
}

func (a *App) Close() error {

	return nil
}
