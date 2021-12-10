package main

import (
	defaultApp "github.com/daqnext/cli-config-template/apps/default_app"
	logApp "github.com/daqnext/cli-config-template/apps/logs_app"
	serviceApp "github.com/daqnext/cli-config-template/apps/service_app"
	"github.com/daqnext/cli-config-template/cli"
)

func main() {
	if cli.AppToDO == nil {
		return
	}

	switch cli.AppToDO.AppName {
	case cli.APP_NAME_LOG:
		logApp.StartLog(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	case cli.APP_NAME_SERVICE:
		serviceApp.RunServiceCmd(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	default:
		cli.LocalLogger.Infoln("======== start default app ===")
		defaultApp.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
