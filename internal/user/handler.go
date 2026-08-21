package uhandl

import (
	"Service/internal/tg-bot"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	tService *tServ.TGService
}

func New(tService *tServ.TGService) *UserHandler {
	return &UserHandler{
		tService: tService,
	}
}

func (h *UserHandler) Register(e *echo.Echo) {
	e.POST("/api/tg", h.Send)
}

func (h *UserHandler) Send(c *echo.Context) error {
	token := c.FormValue("token")
	chatID := c.FormValue("chat_id")
	message := c.FormValue("message")

	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return c.JSON(400, "invalid chat_id")
	}

	if err := h.tService.Send(token, id, message); err != nil {
		return err
	}

	return nil
}
