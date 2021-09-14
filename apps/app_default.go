package apps

import (
	"fmt"

	"github.com/daqnext/cli-config-template/utils"
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

func StartDefault(ConfigJson *fj.FastJson, CliContext *cli.Context) {
	fmt.Println(string(utils.Purple), "hello world , this default app")

}

func init() {

}
