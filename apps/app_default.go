package apps

import (
	"fmt"

	"github.com/daqnext/cli-config-template/global"
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

func StartDefault(ConfigJson *fj.FastJson, CliContext *cli.Context) {
	fmt.Println("this is the default app")
	fmt.Println(ConfigJson.GetContentAsString())
	//fmt.Println(CliContext)
	fmt.Println(global.Something)
	fmt.Println("end of app")
}

func init() {
	//fmt.Println("app_default.go init")
}
