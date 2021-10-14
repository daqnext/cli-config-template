package defaultapp

import (
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
	global.EchoServer.Start()
}

func init() {

}
