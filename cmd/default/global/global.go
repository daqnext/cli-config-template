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
	elasticsearch "github.com/olivere/elastic/v7"
	"gorm.io/gorm"

	"github.com/daqnext/ESUploader/uploader"
)

var ESUploader *uploader.Uploader
var Redis *redis.ClusterClient
var EchoServer *components.EchoServer
var SpMgr *SPR_go.SprJobMgr
var BGJobM *bgjob.JobManager
var sqlDB *sql.DB
var GormDB *gorm.DB
var LocalCache *gofastcache.LocalCache
var EctServer *server.EctHttpServer
var InfuraClient *components.InfuraClient
var ElasticSClient *elasticsearch.Client

func Init() {
	var err error
	//init your global components
	LocalCache = components.InitFastCache(cli.LocalLogger)
	components.InitSmartRoutine()
	BGJobM = components.InitBGJobs(cli.LocalLogger)

	cli.LocalLogger.Info("init system .....")
	////////////ini more components config as you need///////////////////

	EchoServer, err = components.InitEchoServer(cli.LocalLogger, cli.CmdToDo.ConfigJson)
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}

	// GormDB, sqlDB, err = components.InitDB(cli.LocalLogger, cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	// Redis, err = components.InitRedis(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }
	// SpMgr, err = components.InitSprJobs(cli.LocalLogger, cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }
	// EctServer, err = components.InitEctmServer(cli.LocalLogger, cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	// InfuraClient, err = components.InitInfura(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	// ElasticSClient, err = components.InitElasticSearch(cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	// ESUploader, err = components.InitESUploader(cli.LocalLogger, cli.AppToDO.ConfigJson)
	// if err != nil {
	// 	cli.LocalLogger.Fatal(err.Error())
	// }

	cli.LocalLogger.Info("=========== end of init system ==================")

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
