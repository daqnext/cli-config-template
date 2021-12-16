package cli

import (
	"github.com/daqnext/utils/color_util"
	"log"
	"os"
)

func ReadArgs() {
	//print any initialzation panic
	defer func() {
		if err := recover(); err != nil {
			color_util.ColorPrintln(color_util.Red, "panic errors:", err.(error).Error())
		}
	}()

	//ini logger
	IniLocalLogger()

	//config app to run
	errRun := configCliCmd().Run(os.Args)
	if errRun != nil {
		log.Fatal(errRun)
		panic(errRun.Error())
	}
}
