package main

import (
	defaultApp "github.com/daqnext/cli-config-template/apps/default_app"
	logApp "github.com/daqnext/cli-config-template/apps/logs_app"
	"github.com/daqnext/cli-config-template/cli"
)

func main() {

	switch cli.AppToDO.AppName {
	case cli.APP_LOG_NAME:
		logApp.StartLog(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
		break
	default:
		cli.LocalLogger.Infoln("======== start default app ===")
		defaultApp.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
