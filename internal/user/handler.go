package user

import (
	"Service/internal/models"
	"Service/internal/tg-bot"
	"strconv"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	tService   *tServ.TGService
	repository *Repository
}

func New(tService *tServ.TGService, repo *Repository) *UserHandler {
	return &UserHandler{
		tService:   tService,
		repository: repo,
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

	ctx := c.Request().Context()

	req := models.Request{
		Token:   token,
		Chat_id: id,
		Message: message,
		Ip:      c.RealIP(),
	}

	if err := h.tService.Send(token, id, message); err != nil {
		return err
	}

	if err := h.repository.StoreData(ctx, req); err != nil {
		return err
	}

	return nil
}
