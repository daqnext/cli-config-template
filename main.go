package main

import (
	defaultApp "github.com/daqnext/cli-config-template/apps/default"
	"github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/global"
)

func main() {

	switch cli.AppToDO.AppName {
	default:
		global.LocalLogger.Infoln("======== start default app ===")
		defaultApp.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
