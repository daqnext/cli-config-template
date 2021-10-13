package somepack

import (
	"github.com/daqnext/cli-config-template/cli"
)

func HowToGetGlobalParam() {
	cli.LocalLogger.Infoln(cli.AppToDO.ConfigJson.GetString("config"))
}
