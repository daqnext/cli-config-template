package apps

import (
	"strconv"

	_ "github.com/daqnext/cli-config-template/apps/default/controllers"
	"github.com/daqnext/cli-config-template/apps/default/somepack"
	"github.com/daqnext/cli-config-template/global"
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

func StartDefault(ConfigJson *fj.FastJson, CliContext *cli.Context) {

	defer func() {
		global.ReleaseResource()
	}()

	global.LocalLogger.Infoln("hello world , this default app")

	somepack.HowToGetGlobalParam()
	///start the server
	http_port, err := ConfigJson.GetInt("http_port")
	if err != nil {
		global.LocalLogger.Fatalln("http_port is not configured")
	}
	global.LocalLogger.Infoln("http server echo started on port:", http_port)
	global.Echo.Start(":" + strconv.Itoa(http_port))
}

func init() {

}
