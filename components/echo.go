package components

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	localLog "github.com/daqnext/LocalLog/log"
	fj "github.com/daqnext/fastjson"

	"github.com/daqnext/utils/path_util"
	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	Echo                   *echo.Echo
	Http_port              int
	Http_static_abs_folder string
	localLogger            *localLog.LocalLog
}

/*
http_port
http_static_rel_folder
*/
func InitEchoServer(localLogger_ *localLog.LocalLog, ConfigJson *fj.FastJson) (*EchoServer, error) {

	http_port, err := ConfigJson.GetInt("http_port")
	if err != nil {
		return nil, errors.New("http_port [int] in config.json not defined," + err.Error())
	}

	http_static_rel_folder, err := ConfigJson.GetString("http_static_rel_folder")
	if err != nil {
		return nil, errors.New("http_static_rel_folder [string] in config.json not defined," + err.Error())
	}

	esP := &EchoServer{
		echo.New(),
		http_port,
		path_util.GetAbsPath(http_static_rel_folder),
		localLogger_,
	}

	//set locallogger
	esP.Echo.Use(NewEchoLogger(localLogger_))
	return esP, nil
}

func (s *EchoServer) Start() {
	s.localLogger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
	s.localLogger.Infoln("http server with static folder:" + s.Http_static_abs_folder)
	s.Echo.Static("/", s.Http_static_abs_folder)
	s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}

func NewEchoLogger(l *localLog.LocalLog) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
				//error log
				l.WithFields(localLog.Fields{
					"request":     c.Request().RequestURI,
					"method":      c.Request().Method,
					"remote":      c.Request().RemoteAddr,
					"status":      c.Response().Status,
					"text_status": http.StatusText(c.Response().Status),
					"took":        time.Since(start),
					"request_id":  c.Request().Header.Get("X-Request-Id"),
				}).Errorln("request error:" + err.Error())

			} else {
				//info log ,only for loglevels : debug or trace
				if l.Level >= localLog.LLEVEL_DEBUG {
					lentry := l.WithFields(localLog.Fields{
						"request":     c.Request().RequestURI,
						"method":      c.Request().Method,
						"remote":      c.Request().RemoteAddr,
						"status":      c.Response().Status,
						"text_status": http.StatusText(c.Response().Status),
						"took":        time.Since(start),
						"request_id":  c.Request().Header.Get("X-Request-Id"),
					})
					if l.Level == localLog.LLEVEL_DEBUG {
						lentry.Infoln("request success")
					} else {
						lentry.Traceln("request success")
					}

				}
			}
			return nil
		}

	}
}
