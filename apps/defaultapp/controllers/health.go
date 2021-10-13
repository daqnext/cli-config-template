package controllers

import (
	"time"

	"github.com/daqnext/cli-config-template/global"
	"github.com/labstack/echo/v4"
)

func init() {

	global.Echo.GET("/heartbeat", func(c echo.Context) error {
		r := struct {
			UnixTime int64
		}{
			UnixTime: time.Now().Unix(),
		}
		return c.JSON(200, r)
	})
}
