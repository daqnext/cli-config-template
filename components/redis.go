package components

import (
	"context"
	"errors"
	"strconv"

	fj "github.com/daqnext/fastjson"
	"github.com/go-redis/redis/v8"
)

/*
redis_addr
redis_username
redis_password
redis_port
*/

func InitRedis(ConfigJson *fj.FastJson) (*redis.ClusterClient, error) {

	redis_addr, _redis_addr_err := ConfigJson.GetString("redis_addr")
	if _redis_addr_err != nil {
		return nil, errors.New("redis_addr [string] in config.json not defined," + _redis_addr_err.Error())
	}

	redis_username, redis_username_err := ConfigJson.GetString("redis_username")
	if redis_username_err != nil {
		return nil, errors.New("redis_username [string] in config.json not defined," + redis_username_err.Error())
	}

	redis_password, redis_password_err := ConfigJson.GetString("redis_password")
	if redis_password_err != nil {
		return nil, errors.New("redis_password [string] in config.json not defined," + redis_password_err.Error())
	}

	redis_port, redis_port_err := ConfigJson.GetInt("redis_port")
	if redis_port_err != nil {
		return nil, errors.New("redis_port [int] in config.json not defined," + redis_port_err.Error())
	}

	Redis := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redis_addr + ":" + strconv.Itoa(redis_port)},
		Username: redis_username,
		Password: redis_password,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return Redis, nil
}
