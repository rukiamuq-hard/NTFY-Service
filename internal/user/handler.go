package uhandl

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
}

func New() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(e *echo.Echo) {
	e.GET("/user", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, c.RealIP())
	})
}
