package apps

import (
	"fmt"

	_ "github.com/daqnext/cli-config-template/apps/default/controllers"
	"github.com/daqnext/cli-config-template/apps/default/somepack"
	"github.com/daqnext/cli-config-template/global"
	"github.com/daqnext/cli-config-template/utils"
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

func StartDefault(ConfigJson *fj.FastJson, CliContext *cli.Context) {
	fmt.Println(string(utils.Purple), "hello world , this default app")

	somepack.HowToGetGlobalParam()

	///start the server
	http_port, err := ConfigJson.GetString("http_port")
	if err != nil {
		panic("http_port is not configured")
	}
	global.Echo.Start(":" + http_port)

}

func init() {

}
