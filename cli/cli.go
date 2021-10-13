package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/daqnext/cli-config-template/utils"
)

func init() {
	//print any initialzation panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(utils.Red), "panic errors:", err.(error).Error())
		}
	}()

	//config exe path
	configAbsPath()
	//ini logger
	iniLocalLogger()
	//config app to run
	errRun := configCliApp().Run(os.Args)
	if errRun != nil {
		log.Fatal(errRun)
		panic(errRun.Error())
	}
}
