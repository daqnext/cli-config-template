package main

import (
	"fmt"

	defaultApp "github.com/daqnext/cli-config-template/apps/default"
	"github.com/daqnext/cli-config-template/cli"
	_ "github.com/daqnext/cli-config-template/global"
	"github.com/daqnext/cli-config-template/utils"
)

func main() {

	switch cli.AppToDO.AppName {
	default:
		fmt.Println(string(utils.Green), "======== start default app ===")
		defaultApp.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
