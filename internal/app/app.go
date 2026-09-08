package app

import (
	ds "Service/internal/discord"   // discord service
	es "Service/internal/echo-serv" // echo server
	"Service/internal/notification" // ntfy service
	"Service/internal/repository"   // repository
	ts "Service/internal/telegram"  // telegram service
)

type App struct {
	serv        *es.Server
	uhandler    *notification.NotifyHandler
	tservice    *ts.TGService
	dservice    *ds.DSService
	mrepository *repository.Postgress // main repository
	urepository *notification.Repository
}

func New() *App {
	return &App{}
}

func (a *App) Start() error {

	a.mrepository = repository.New()
	if err := a.mrepository.Connect(); err != nil {
		return err
	}

	a.urepository = notification.NewUserRepo(a.mrepository.DB)

	a.serv = es.New()     // echo server
	a.tservice = ts.New() // tg service
	a.dservice = ds.New() // ds service

	//handlers
	a.uhandler = notification.New(a.tservice, a.dservice, a.urepository)
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
