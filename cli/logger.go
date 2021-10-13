package cli

import (
	"fmt"

	localLog "github.com/daqnext/LocalLog/log"
	"github.com/daqnext/cli-config-template/utils"
)

var LocalLogger *localLog.LocalLog

func iniLocalLogger() {
	var llerr error
	LocalLogger, llerr = localLog.New(GetPath("logs"), 2, 20, 30)

	if llerr != nil {
		fmt.Println(string(utils.Red), "Error:")
		panic("local_log error:" + llerr.Error())
	}
}
