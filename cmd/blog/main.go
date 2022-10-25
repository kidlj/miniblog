package main

import (
	"github.com/kidlj/blog/blog"
	"github.com/kidlj/blog/dist"
	"github.com/kidlj/blog/index"
	"github.com/kidlj/blog/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	t := templates.NewTemplate()
	e.Renderer = t

	indexHandler := index.NewHandler()
	indexHandler.InstallRoutes(e)

	staticHandler := dist.NewHandler()
	staticHandler.InstallRoutes(e)

	blogService := blog.NewService()
	blogHandler := blog.NewHandler(blogService)
	blogHandler.InstallRoutes(e)

	err := e.Start(blog.BLOG_LISTEN_ADDR)
	if err != nil {
		panic(err)
	}
}
