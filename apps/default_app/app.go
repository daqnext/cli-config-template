package defaultapp

import (
	"strconv"

	_ "github.com/daqnext/cli-config-template/apps/default_app/controllers"
	"github.com/daqnext/cli-config-template/apps/default_app/global"
	"github.com/daqnext/cli-config-template/apps/default_app/somepack"
	clitool "github.com/daqnext/cli-config-template/cli"
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

func StartDefault(ConfigJson *fj.FastJson, CliContext *cli.Context) {

	defer func() {
		global.ReleaseResource()
	}()

	//print logo
	clitool.LocalLogger.Infoln(clitool.Logo)
	clitool.LocalLogger.Infoln("hello world , this default app")
	somepack.HowToGetGlobalParam()
	///start the server
	http_port, err := ConfigJson.GetInt("http_port")
	if err != nil {
		clitool.LocalLogger.Fatalln("http_port is not configured")
	}
	clitool.LocalLogger.Infoln("http server echo started on port:", http_port)
	global.Echo.Start(":" + strconv.Itoa(http_port))
}

func init() {

}
