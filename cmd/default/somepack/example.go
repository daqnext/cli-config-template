package somepack

import (
	"github.com/daqnext/cli-config-template/cli"
)

func HowToGetGlobalParam() {
	cli.LocalLogger.Infoln(cli.CmdToDo.ConfigJson.GetString("config"))
}
