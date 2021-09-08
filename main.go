package main

import (
	"github.com/daqnext/cli-config-template/apps"
	"github.com/daqnext/cli-config-template/cli"
	_ "github.com/daqnext/cli-config-template/global"
)

func main() {

	switch cli.AppToDO.AppName {
	default:
		apps.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
