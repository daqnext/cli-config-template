package controllers

import (
	"time"

	"github.com/daqnext/cli-config-template/apps/default_app/global"
	"github.com/labstack/echo/v4"
)

func init() {
	if !global.GLOBAL_INIT_FINISHED {
		return
	}

	global.EchoServer.Echo.GET("/heartbeat", func(c echo.Context) error {
		r := struct {
			UnixTime int64
		}{
			UnixTime: time.Now().Unix(),
		}
		return c.JSON(200, r)
	})
}
