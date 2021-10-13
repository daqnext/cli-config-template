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

var Redis *redis.ClusterClient
var Echo *echo.Echo

var SpMgr *SPR_go.SprJobMgr
var BGJobM *bgjob.JobManager
var sqlDB *sql.DB
var GormDB *gorm.DB
var LocalCache *gofastcache.LocalCache

var ExEPath string

func init() {
	if !cli.AppIsActive(cli.APP_DEFAULT_NAME) {
		return
	}

	//first step to init log
	initLocalLog()
	//init your global components
	SmartRoutine.ClearPanics()
	LocalCache = gofastcache.New()
	//initDB()
	//initRedis()
	//iniGJobs()
	//initSprJobs() //only for server-side app
	Echo = echo.New()

}

func initLocalLog() {

	local_log_level, local_log_level_err := cli.AppToDO.ConfigJson.GetString("local_log_level")
	if local_log_level_err != nil {
		cli.LocalLogger.Fatalln("local_log_level not configured")
	}
	cli.LocalLogger.ResetLevel(local_log_level)
}

func initRedis() {

	redis_addr, _redis_addr_err := cli.AppToDO.ConfigJson.GetString("redis_addr")
	if _redis_addr_err != nil {
		cli.LocalLogger.Fatalln("redis_addr not configured")
	}

	redis_port, redis_port_err := cli.AppToDO.ConfigJson.GetInt("redis_port")
	if redis_port_err != nil {
		cli.LocalLogger.Fatalln("redis_port not configured")
	}

	redis_username, redis_username_err := cli.AppToDO.ConfigJson.GetString("redis_username")
	if redis_username_err != nil {
		cli.LocalLogger.Fatalln("redis_username not configured")
	}

	redis_password, redis_password_err := cli.AppToDO.ConfigJson.GetString("redis_password")
	if redis_password_err != nil {
		cli.LocalLogger.Fatalln("redis_password not configured")
	}

	Redis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redis_addr + ":" + strconv.Itoa(redis_port)},
		Username: redis_username,
		Password: redis_password,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		cli.LocalLogger.Fatalln("Redis connect failed")
	}

}

func iniGJobs() {
	//////// ini bGJob   //////////////////////
	BGJobM = bgjob.New()
}

//only for server-side app
func initSprJobs() {

	//////// ini spr job //////////////////////
	redis_addr, _redis_addr_err := cli.AppToDO.ConfigJson.GetString("redis_addr")
	if _redis_addr_err != nil {
		cli.LocalLogger.Fatalln("redis_addr not configured")
	}

	redis_port, redis_port_err := cli.AppToDO.ConfigJson.GetInt("redis_port")
	if redis_port_err != nil {
		cli.LocalLogger.Fatalln("redis_port not configured")
	}

	redis_username, redis_username_err := cli.AppToDO.ConfigJson.GetString("redis_username")
	if redis_username_err != nil {
		cli.LocalLogger.Fatalln("redis_username not configured")
	}

	redis_password, redis_password_err := cli.AppToDO.ConfigJson.GetString("redis_password")
	if redis_password_err != nil {
		cli.LocalLogger.Fatalln("redis_password not configured")
	}

	var SPR_go_err error
	SpMgr, SPR_go_err = SPR_go.New(SPR_go.RedisConfig{
		Addr:     redis_addr,
		Port:     redis_port,
		UserName: redis_username,
		Password: redis_password,
	})
	if SPR_go_err != nil {
		panic(SPR_go_err.Error())
	}

}

func initDB() {

	db_host, db_host_err := cli.AppToDO.ConfigJson.GetString("db_host")
	if db_host_err != nil {
		cli.LocalLogger.Fatalln("db_host not configured")
	}

	db_port, db_port_err := cli.AppToDO.ConfigJson.GetInt("db_port")
	if db_port_err != nil {

		cli.LocalLogger.Fatalln("db_port not configured")
	}

	db_name, db_name_err := cli.AppToDO.ConfigJson.GetString("db_name")
	if db_name_err != nil {
		cli.LocalLogger.Fatalln("db_name not configured")
	}

	db_username, db_username_err := cli.AppToDO.ConfigJson.GetString("db_username")
	if db_username_err != nil {
		cli.LocalLogger.Fatalln("db_username not configured")
	}

	db_password, db_password_err := cli.AppToDO.ConfigJson.GetString("db_password")
	if db_password_err != nil {
		cli.LocalLogger.Fatalln("db_password not configured")
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	var erropen error
	GormDB, erropen = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//some config
	})
	if erropen != nil {
		cli.LocalLogger.Fatalln("failed to connect databas")
	}
	//set pool
	var sqlerr error
	sqlDB, sqlerr = GormDB.DB()
	if sqlerr != nil {
		cli.LocalLogger.Fatalln("failed to get database")
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
