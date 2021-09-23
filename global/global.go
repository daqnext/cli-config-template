package global

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/daqnext/BGJOB_GO/bgjob"
	SPR_go "github.com/daqnext/SPR-go"
	"github.com/daqnext/cli-config-template/cli"
	gofastcache "github.com/daqnext/go-fast-cache"
	SmartRoutine "github.com/daqnext/go-smart-routine/sr"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

///declear the global components
var Redis *redis.Client
var Echo *echo.Echo

var SpMgr *SPR_go.SprJobMgr
var BGJobM *bgjob.JobManager
var sqlDB *sql.DB
var LocalCache *gofastcache.LocalCache

func init() {

	//init your global components
	SmartRoutine.ClearPanics()
	LocalCache = gofastcache.New()
	//initDB()
	//initRedis()
	//iniGJobs()
	//initSprJobs()
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

func iniGJobs() {
	//////// ini bGJob   //////////////////////
	BGJobM = bgjob.New()
}

func initSprJobs() {

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

func initDB() {

	db_host, db_host_err := cli.AppToDO.ConfigJson.GetString("db_host")
	if db_host_err != nil {
		panic("db_host not configured")
	}

	db_port, db_port_err := cli.AppToDO.ConfigJson.GetString("db_port")
	if db_port_err != nil {
		panic("db_port not configured")
	}

	db_name, db_name_err := cli.AppToDO.ConfigJson.GetString("db_name")
	if db_name_err != nil {
		panic("db_name not configured")
	}

	db_username, db_username_err := cli.AppToDO.ConfigJson.GetString("db_username")
	if db_username_err != nil {
		panic("db_username not configured")
	}

	db_password, db_password_err := cli.AppToDO.ConfigJson.GetString("db_password")
	if db_password_err != nil {
		panic("db_password not configured")
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=UTC"

	dbc, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//some config
	})
	if err != nil {
		panic("failed to connect database")
	}
	//设置数据库连接池
	sqlDB, err = dbc.DB()
	if err != nil {
		panic("failed to get database")
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

}

func ReleaseResource() {
	if Redis != nil {
		Redis.Close()
	}
	if sqlDB != nil {
		sqlDB.Close()
	}
	if Echo != nil {
		Echo.Close()
	}
}
