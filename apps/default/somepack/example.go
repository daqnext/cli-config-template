package somepack

import (
	"github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/global"
)

func HowToGetGlobalParam() {
	global.LocalLogger.Infoln(cli.AppToDO.ConfigJson.GetString("config"))
}
