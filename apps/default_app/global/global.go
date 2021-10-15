package global

import (
	"database/sql"

	"github.com/daqnext/BGJOB_GO/bgjob"
	"github.com/daqnext/ECTSM-go/http/server"
	SPR_go "github.com/daqnext/SPR-go"
	elasticsearch "github.com/olivere/elastic/v7"

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
var ESClient *elasticsearch.Client

func init() {

	if !cli.AppIsActive(cli.APP_NAME_DEFAULT) {
		return
	}

	//first step to init log
	inilogerr := components.InitLocalLog(cli.LocalLogger, cli.AppToDO.ConfigJson)
	if inilogerr != nil {
		panic(inilogerr.Error())
	}
	components.InitSmartRoutine()
	BGJobM = components.InitBGJobs()
	LocalCache = components.InitFastCache()

	var err error
	EchoServer, err = components.InitEchoServer(cli.LocalLogger, cli.AppToDO.ConfigJson)
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}
	////////////ini more components config as you need///////////////////
	// GormDB, sqlDB, err = components.InitDB(cli.LocalLogger, cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }
	// Redis, err = components.InitRedis(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }
	// SpMgr, err = components.InitSprJobs(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }
	// EctServer, err = components.InitEctmServer(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	// InfuraClient, err = components.InitInfura(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	// ESClient, err = components.InitElasticSearch(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

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
