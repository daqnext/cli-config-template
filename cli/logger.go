package cli

import (
	localLog "github.com/daqnext/LocalLog/log"
	"github.com/daqnext/utils/color_util"
	"github.com/daqnext/utils/path_util"
)

var LocalLogger *localLog.LocalLog

func iniLocalLogger() {

	var llerr error
	LocalLogger, llerr = localLog.New(path_util.GetAbsPath("logs"), 2, 20, 30)

	if llerr != nil {
		color_util.ColorPrintln(color_util.Red, "Error:")
		panic("local_log error:" + llerr.Error())
	}
}
