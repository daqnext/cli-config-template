package defaultapp

import (
	clitool "github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/cmd/default/controllers"
	_ "github.com/daqnext/cli-config-template/cmd/default/controllers"
	"github.com/daqnext/cli-config-template/cmd/default/global"
	"github.com/daqnext/cli-config-template/cmd/default/somepack"
)

func StartDefault() {
	global.Init()
	defer func() {
		clitool.LocalLogger.Infoln("StartDefault closed , start to ReleaseResource()")
		global.ReleaseResource()
	}()

	controllers.DeployApi()

	//print logo
	clitool.LocalLogger.Infoln(clitool.Logo)
	clitool.LocalLogger.Infoln("hello world , this default app")
	somepack.HowToGetGlobalParam()
	///start the server
	err := global.EchoServer.Start()
	if err != nil {
		clitool.LocalLogger.Fatalln(err)
	}

}
