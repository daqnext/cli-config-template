package main

import (
	"fmt"

	"github.com/daqnext/cli-config-template/apps"
	"github.com/daqnext/cli-config-template/cli"
	_ "github.com/daqnext/cli-config-template/global"
	"github.com/daqnext/cli-config-template/utils"
)

func main() {

	switch cli.AppToDO.AppName {
	default:
		fmt.Println(string(utils.Green), "======== start default app ===")
		apps.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
