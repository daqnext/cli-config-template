package controllers

import (
	"time"

	"github.com/daqnext/cli-config-template/cmd/default/global"
	"github.com/labstack/echo/v4"
)

func healthRouter() {

	global.EchoServer.Echo.GET("/heartbeat", func(c echo.Context) error {
		r := struct {
			UnixTime int64
		}{
			UnixTime: time.Now().Unix(),
		}
		return c.JSON(200, r)
	})
}
