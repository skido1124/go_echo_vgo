package handler

import (
	"github.com/labstack/echo"
)

// PingHandler テスト用 ping
func PingHandler(c echo.Context) error {
	return c.JSON(200, "{'message': 'ping'}")
}
