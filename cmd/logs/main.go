package logs

import (
	"github.com/daqnext/cli-config-template/cli"
)

func StartLog() {
	num := cli.CmdToDo.CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := cli.CmdToDo.CliContext.Bool("onlyerr")
	if onlyerr {
		cli.LocalLogger.PrintLastN_ErrLogs(num)
	} else {
		cli.LocalLogger.PrintLastN_AllLogs(num)
	}
}
