package somepack

import (
	"fmt"

	"github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/global"
)

func HowToGetGlobalParam() {
	fmt.Println(global.Something)
	fmt.Println(cli.AppToDO.ConfigJson.GetString("config"))
}
