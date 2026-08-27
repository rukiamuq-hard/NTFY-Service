package echoserv

import (
	"github.com/labstack/echo/v5"
	//"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	ServEcho *echo.Echo
}

func New() *Server {
	e := echo.New()
	return &Server{ServEcho: e}
}

func (s *Server) Start(addr string) error {
	return s.ServEcho.Start(addr)
}
