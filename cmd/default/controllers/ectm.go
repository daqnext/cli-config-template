package controllers

import (
	"time"

	"github.com/daqnext/ECTSM-go/http/server"
	ectsmUtil "github.com/daqnext/ECTSM-go/utils"
	"github.com/daqnext/cli-config-template/cmd/default/global"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ectmRouter() {

	global.EchoServer.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders: []string{
			"ectm_key", "Ectm_key",
			"ectm_time", "Ectm_time",
			"ectm_token", "Ectm_time",
		},
	}))

	global.EchoServer.Echo.GET("/ectminfo", func(c echo.Context) error {
		r := struct {
			UnixTime  int64
			PublicKey string
		}{
			time.Now().Unix(),
			ectsmUtil.PublicKeyToString(&global.EctServer.PrivateKey.PublicKey),
		}
		return c.JSON(200, r)
	})

	global.EchoServer.Echo.GET("/testget", func(c echo.Context) error {
		ECTReq := global.EctServer.HandleGet(c.Request())
		if ECTReq.Err != nil {
			return c.String(500, ECTReq.Err.Error())
		}
		resData := struct {
			Status     int
			Msg        string
			ClientTime int64
		}{0, "get success", time.Now().Unix()}

		sendData, err := server.ECTSendBack(c.Response().Header(), ECTReq.SymmetricKey, resData)
		if err != nil {
			return c.String(500, err.Error())
		}
		return c.String(200, string(sendData))
	})

	global.EchoServer.Echo.POST("/testpost", func(c echo.Context) error {

		ectRequest := global.EctServer.HandlePost(c.Request())
		if ectRequest.Err != nil {
			return c.String(500, "decrypt error:")
		}

		//responseData example
		data := struct {
			Status int
			Msg    string
			Data   interface{}
		}{0, "post success", "this is something"}

		sendData, err := server.ECTSendBack(c.Response().Header(), ectRequest.SymmetricKey, data)
		if err != nil {
			return c.String(500, err.Error())
		}

		return c.Blob(200, "application/octet-stream", sendData)

	})

}
