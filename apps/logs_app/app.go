package logs

import (
	clitool "github.com/daqnext/cli-config-template/cli"
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

func StartLog(ConfigJson *fj.FastJson, CliContext *cli.Context) {

	num := CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := CliContext.Bool("onlyerr")
	if onlyerr {
		clitool.LocalLogger.PrintLastN_ErrLogs(num)
	} else {
		clitool.LocalLogger.PrintLastN_AllLogs(num)
	}
}

func init() {

	if !clitool.AppIsActive(clitool.APP_NAME_LOG) {
		return
	}

}
