package blog

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type Handler struct {
	service *Service
}

func (h *Handler) InstallRoutes(router *echo.Echo) {
	g := router.Group("/blog")

	g.GET("/", h.list)
	g.GET("/:path", h.get)
	g.GET("/feed", h.feed)
}

func (h *Handler) list(c echo.Context) error {
	blogs, err := h.service.list()
	if err != nil {
		return echo.ErrNotFound
	}

	return c.Render(http.StatusOK, "blogs.html", echo.Map{
		"Blogs": blogs,
		"Feed":  BLOG_FEED,
	})
}

func (h *Handler) get(c echo.Context) error {
	path := c.Param("path") + ".md"
	blog, err := h.service.get(path)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.Render(http.StatusOK, "blog.html", echo.Map{
		"Blog": blog,
		"Feed": BLOG_FEED,
	})
}

func (h *Handler) feed(c echo.Context) error {
	feed, err := h.service.feed()
	if err != nil {
		return echo.ErrNotFound
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/atom+xml")
	return c.String(http.StatusOK, feed)
}
