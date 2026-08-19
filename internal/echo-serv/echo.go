package echoserv

import (
	"github.com/labstack/echo/v5"
	//"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	e := echo.New()
	return &Server{echo: e}
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}
