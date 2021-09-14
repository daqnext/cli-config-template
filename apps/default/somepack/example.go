package somepack

import (
	"fmt"

	"github.com/daqnext/cli-config-template/cli"
)

func HowToGetGlobalParam() {
	fmt.Println(cli.AppToDO.ConfigJson.GetString("config"))
}
