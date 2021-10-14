package components

import (
	"errors"
	"strconv"

	"github.com/daqnext/cli-config-template/cli"
	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	Echo                   *echo.Echo
	Http_port              int
	Http_static_abs_folder string
}

/*
http_port
http_static_rel_folder
*/
func InitEchoServer() (*EchoServer, error) {

	http_port, err := cli.AppToDO.ConfigJson.GetInt("http_port")
	if err != nil {
		return nil, errors.New("http_port [int] in config.json not defined," + err.Error())
	}

	http_static_rel_folder, err := cli.AppToDO.ConfigJson.GetString("http_static_rel_folder")
	if err != nil {
		return nil, errors.New("http_static_rel_folder [string] in config.json not defined," + err.Error())
	}

	return &EchoServer{
		echo.New(),
		http_port,
		cli.GetPath(http_static_rel_folder),
	}, nil
}

func (s *EchoServer) Start() {
	cli.LocalLogger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
	cli.LocalLogger.Infoln("http server with static folder:" + s.Http_static_abs_folder)
	s.Echo.Static("/", s.Http_static_abs_folder)
	s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}
