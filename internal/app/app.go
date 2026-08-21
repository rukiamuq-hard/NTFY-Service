package app

import (
	es "Service/internal/echo-serv"
	ts "Service/internal/tg-bot"
	"Service/internal/user"
)

type App struct {
	serv     *es.Server
	uhandler *uhandl.UserHandler
	tservice *ts.TGService
}

func New() *App {
	return &App{}
}

func (a *App) Start() error {
	a.serv = es.New()     // echo server
	a.tservice = ts.New() // tg service

	//handlers
	a.uhandler = uhandl.New(a.tservice)
	a.uhandler.Register(a.serv.ServEcho)

	if err := a.serv.Start(":8080"); err != nil {
		return err
	}

	return nil
}

func (a *App) Close() error {

	return nil
}
