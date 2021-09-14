package global

import (
	"context"
	"strconv"

	"github.com/daqnext/BGJOB_GO/bgjob"
	SPR_go "github.com/daqnext/SPR-go"
	"github.com/daqnext/cli-config-template/cli"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

///declear the global components
var Redis *redis.Client
var Echo *echo.Echo

var SpMgr *SPR_go.SprJobMgr
var BGJobM *bgjob.JobManager

func init() {

	//init your global components

	//initRedis()
	//initJobs()

	Echo = echo.New()

}

func initRedis() {

	redis_addr, _redis_addr_err := cli.AppToDO.ConfigJson.GetString("redis_addr")
	if _redis_addr_err != nil {
		panic("redis_addr not configured")
	}

	redis_port, redis_port_err := cli.AppToDO.ConfigJson.GetInt("redis_port")
	if redis_port_err != nil {
		panic("redis_port not configured")
	}

	redis_db, redis_db_err := cli.AppToDO.ConfigJson.GetInt("redis_db")
	if redis_db_err != nil {
		panic("redis_db not configured")
	}

	Redis = redis.NewClient(&redis.Options{
		Addr: redis_addr + ":" + strconv.FormatInt(redis_port, 10),
		DB:   int(redis_db),
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		panic("Redis connect failed")
	}

}

func initJobs() {

	//////// ini bGJob   //////////////////////
	BGJobM = bgjob.New()

	//////// ini spr job //////////////////////
	redis_addr, _redis_addr_err := cli.AppToDO.ConfigJson.GetString("redis_addr")
	if _redis_addr_err != nil {
		panic("redis_addr not configured")
	}

	redis_port, redis_port_err := cli.AppToDO.ConfigJson.GetInt("redis_port")
	if redis_port_err != nil {
		panic("redis_port not configured")
	}

	redis_db, redis_db_err := cli.AppToDO.ConfigJson.GetInt("redis_db")
	if redis_db_err != nil {
		panic("redis_db not configured")
	}

	var SPR_go_err error
	SpMgr, SPR_go_err = SPR_go.New(SPR_go.RedisConfig{
		Addr: redis_addr,
		Port: int(redis_port),
		Db:   int(redis_db),
	})
	if SPR_go_err != nil {
		panic(SPR_go_err.Error())
	}

}
