package global

import (
	"database/sql"

	"github.com/daqnext/BGJOB_GO/bgjob"
	"github.com/daqnext/ECTSM-go/http/server"
	SPR_go "github.com/daqnext/SPR-go"

	"github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/components"
	gofastcache "github.com/daqnext/go-fast-cache"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

///declear the global components
var Redis *redis.ClusterClient
var EchoServer *components.EchoServer
var SpMgr *SPR_go.SprJobMgr
var BGJobM *bgjob.JobManager
var sqlDB *sql.DB
var GormDB *gorm.DB
var LocalCache *gofastcache.LocalCache
var EctServer *server.EctHttpServer
var InfuraClient *components.InfuraClient

func init() {
	if !cli.AppIsActive(cli.APP_NAME_DEFAULT) {
		return
	}

	//first step to init log
	var err error
	components.InitLocalLog()
	components.InitSmartRoutine()
	BGJobM = components.InitBGJobs()
	LocalCache = components.InitFastCache()

	////////////////////////////////
	GormDB, sqlDB, err = components.InitDB()
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}
	Redis, err = components.InitRedis()
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}
	SpMgr, err = components.InitSprJobs()
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}
	EchoServer, err = components.InitEchoServer()
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}
	EctServer, err = components.InitEctmServer()
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}

	InfuraClient, err = components.InitInfura()
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}

}

func ReleaseResource() {
	if Redis != nil {
		Redis.Close()
	}
	if sqlDB != nil {
		sqlDB.Close()
	}
	if EchoServer != nil {
		EchoServer.Close()
	}
}
