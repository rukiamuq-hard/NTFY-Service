package notification

import (
	"Service/internal/discord"
	"Service/internal/models"
	"Service/internal/telegram"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type NotifyHandler struct {
	tService   *tServ.TGService
	dService   *discord.DSService
	repository *Repository
}

func New(tService *tServ.TGService, dService *discord.DSService, repo *Repository) *NotifyHandler {
	return &NotifyHandler{
		tService:   tService,
		dService:   dService,
		repository: repo,
	}
}

func (h *NotifyHandler) Register(e *echo.Echo) {
	e.POST("/api/tg", h.SendTelegram)
	e.POST("/api/ds", h.SendDiscord)
}

func (h *NotifyHandler) SendTelegram(c *echo.Context) error {
	var hook models.RequestTelegram
	if err := c.Bind(&hook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Status": "Incorrect POST request"})
	}

	ctx := c.Request().Context()

	if err := h.tService.Send(hook); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Status": "Failed to sent message!"})
	}

	repohook := models.NotificationLog{
		Recipient: strconv.FormatInt(hook.ChatID, 10),
		Provider:  "Telegram",
		Message:   hook.Message,
		IP:        c.RealIP(),
	}

	if err := h.repository.StoreData(ctx, repohook); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Status": "Failed to store data!"}) //delete too
	}

	return c.JSON(http.StatusOK, map[string]string{"Status": "Sent!"})
}

func (h *NotifyHandler) SendDiscord(c *echo.Context) error {
	var hook models.NotificationLog
	if err := c.Bind(&hook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Status": "Incorrect POST request"})
	}

	hook.Provider = "Discord"
	hook.IP = c.RealIP()

	ctx := c.Request().Context()

	if err := h.dService.Send(ctx, hook); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Status": "Failed to sent message!"})
	}

	if err := h.repository.StoreData(ctx, hook); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Status": "Failed to store data!"}) //delete too
	}

	return c.JSON(http.StatusOK, map[string]string{"Status": "Sent!"})
}
