package index

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) InstallRoutes(router *echo.Echo) {
	router.GET("/", h.index)
}

func (h *Handler) index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
