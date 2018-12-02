package main

import (
	"github.com/skido1124/go_echo_vgo/handler"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", handler.PingHandler)

	e.Start(":3000")
}
