package main

import (
	"github.com/daqnext/cli-config-template/cli"
	defaultCmd "github.com/daqnext/cli-config-template/cmd/default"
	logCmd "github.com/daqnext/cli-config-template/cmd/logs"
	serviceCmd "github.com/daqnext/cli-config-template/cmd/service"
)

func main() {
	cli.ReadArgs()

	switch cli.CmdToDo.CmdName {
	case cli.CMD_NAME_LOG:
		logCmd.StartLog()
	case cli.CMD_NAME_SERVICE:
		serviceCmd.RunServiceCmd()
	default:
		cli.LocalLogger.Infoln("======== start default app ===")
		defaultCmd.StartDefault()
	}
}
