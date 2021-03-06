package components

import (
	"errors"

	"github.com/daqnext/ECTSM-go/http/server"
	localLog "github.com/daqnext/LocalLog/log"
	fj "github.com/daqnext/fastjson"
)

/*
ectm_s_pubkey
ectm_s_prikey
*/
func InitEctmServer(localLogger_ *localLog.LocalLog, ConfigJson *fj.FastJson) (*server.EctHttpServer, error) {

	_, err_pubkey := ConfigJson.GetString("ectm_s_pubkey")
	if err_pubkey != nil {
		return nil, errors.New("ectm_s_pubkey [string] in config.json not defined," + err_pubkey.Error())

	}

	ECTM_S_PriKey, err_prikey := ConfigJson.GetString("ectm_s_prikey")
	if err_prikey != nil {
		return nil, errors.New("ectm_s_prikey [string] in config.json not defined," + err_prikey.Error())
	}

	ECTM_Server, err_ectmservernew := server.New(ECTM_S_PriKey, localLogger_)
	if err_ectmservernew != nil {
		return nil, err_ectmservernew
	}

	return ECTM_Server, nil

}
