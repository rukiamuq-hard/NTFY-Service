package app

import (
	"Service/internal/echo-serv"
	"Service/internal/user"
)

type App struct {
	serv     *echoserv.Server
	uhandler *uhandl.UserHandler
}

func New() *App {
	return &App{}
}

func (a *App) Start() error {
	a.serv = echoserv.New()

	//handlers
	a.uhandler = uhandl.New()
	a.uhandler.Register(a.serv.ServEcho)

	if err := a.serv.Start(":8080"); err != nil {
		return err
	}

	return nil
}

func (a *App) Close() error {

	return nil
}
