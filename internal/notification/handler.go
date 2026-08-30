package notification

import (
	"Service/internal/models"
	"Service/internal/telegram"
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
	var hook models.WebhookReceiver
	if err := c.Bind(&hook); err != nil {
		return c.JSON(400, map[string]string{"Status": "Incorrect POST request"})
	}
	hook.IP = c.RealIP()

	ctx := c.Request().Context()

	if err := h.tService.Send(hook); err != nil {
		return c.JSON(400, map[string]string{"Status": "Error sending to Telegram!"})
	}

	if err := h.repository.StoreData(ctx, hook); err != nil {
		return c.JSON(400, map[string]string{"Status": "Error with repository!"})
	}

	return c.JSON(200, map[string]string{"Status": "Sended!"})
}
