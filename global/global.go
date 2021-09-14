package global

import (
	_ "github.com/daqnext/cli-config-template/cli"
)

///declear the global components

var Something string

func init() {
	Something = "this is something example"
	//init your global components
	//fmt.Println("global.go init")

}
