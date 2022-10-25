package dist

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed static/*
var static embed.FS

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) InstallRoutes(router *echo.Echo) {
	g := router.Group("/static")
	contentHandler := echo.WrapHandler(http.FileServer(http.FS(static)))
	g.GET("/*", contentHandler)
}
