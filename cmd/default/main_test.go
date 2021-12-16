package defaultapp

import (
	"github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/cli-config-template/cmd/default/controllers"
	"github.com/daqnext/cli-config-template/cmd/default/global"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Init() {
	cli.IniLocalLogger()

	p, _ := os.Getwd()
	cli.ManualInitAppConfig(filepath.Join(p, "..", "..", "configs", "dev.json"))

	global.Init()
	controllers.DeployApi()
}

func Test_appTest(t *testing.T) {
	Init()
	log.Println("app test")
}

func Test_appTest2(t *testing.T) {
	Init()
	log.Println("app test2")
}
