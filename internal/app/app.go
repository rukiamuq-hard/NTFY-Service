package app

import (
	es "Service/internal/echo-serv"
	"Service/internal/repository"
	ts "Service/internal/tg-bot"
	"Service/internal/user"
)

type App struct {
	serv        *es.Server
	uhandler    *user.UserHandler
	tservice    *ts.TGService
	mrepository *repository.Repository // main repository
	urepository *user.Repository
}

func New() *App {
	return &App{}
}

func (a *App) Start() error {

	a.mrepository = repository.New()
	if err := a.mrepository.Connect(); err != nil {
		return err
	}

	a.urepository = user.NewUserRepo(a.mrepository.Mydb)

	a.serv = es.New()     // echo server
	a.tservice = ts.New() // tg service

	//handlers
	a.uhandler = user.New(a.tservice, a.urepository)
	a.uhandler.Register(a.serv.ServEcho)

	if err := a.serv.Start(":8080"); err != nil {
		return err
	}

	return nil
}

func (a *App) Close() error {
	a.mrepository.Close()
	return nil
}
