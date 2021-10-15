package components

import (
	"errors"

	SPR_go "github.com/daqnext/SPR-go"
	fj "github.com/daqnext/fastjson"
)

/*
redis_addr
redis_username
redis_password
redis_port
*/

func InitSprJobs(ConfigJson *fj.FastJson) (*SPR_go.SprJobMgr, error) {

	//////// ini spr job //////////////////////
	redis_addr, _redis_addr_err := ConfigJson.GetString("redis_addr")
	if _redis_addr_err != nil {
		return nil, errors.New("redis_addr [string] in config.json not defined," + _redis_addr_err.Error())
	}

	redis_port, redis_port_err := ConfigJson.GetInt("redis_port")
	if redis_port_err != nil {
		return nil, errors.New("redis_port [int] in config.json not defined," + redis_port_err.Error())
	}

	redis_username, redis_username_err := ConfigJson.GetString("redis_username")
	if redis_username_err != nil {
		return nil, errors.New("redis_username [string] in config.json not defined," + redis_username_err.Error())
	}

	redis_password, redis_password_err := ConfigJson.GetString("redis_password")
	if redis_password_err != nil {
		return nil, errors.New("redis_password [string] in config.json not defined," + redis_password_err.Error())
	}

	SpMgr, SPR_go_err := SPR_go.New(SPR_go.RedisConfig{
		Addr:     redis_addr,
		Port:     redis_port,
		Password: redis_password,
		UserName: redis_username,
	})

	if SPR_go_err != nil {
		return nil, SPR_go_err
	}

	return SpMgr, nil

}
