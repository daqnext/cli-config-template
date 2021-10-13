package main

import (
	defaultApp "github.com/daqnext/cli-config-template/apps/defaultapp"
	logApp "github.com/daqnext/cli-config-template/apps/logs"
	"github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/global"
)

func main() {

	switch cli.AppToDO.AppName {
	case "logs":
		global.LocalLogger.Infoln("======== start logs app ===")
		logApp.StartLog(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
		break
	default:
		global.LocalLogger.Infoln("======== start default app ===")
		defaultApp.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
