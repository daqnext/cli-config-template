package cli

import (
	"github.com/daqnext/utils/color_util"
	"log"
	"os"
)

func init() {
	//if strings.HasSuffix(os.Args[0], ".test") {
	//	return
	//}

	//print any initialzation panic
	defer func() {
		if err := recover(); err != nil {
			color_util.ColorPrintln(color_util.Red, "panic errors:", err.(error).Error())
		}
	}()

	//ini logger
	iniLocalLogger()
	//config app to run
	errRun := configCliApp().Run(os.Args)
	if errRun != nil {
		log.Fatal(errRun)
		panic(errRun.Error())
	}
}
