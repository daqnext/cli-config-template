package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/daqnext/cli-config-template/utils"
)

var ExEPath string

func GetPath(relpath string) string {
	return ExEPath + "/" + strings.Trim(relpath, "/")
}

func configAbsPath() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err.Error())
	}
	runPath, err := filepath.Abs(file)
	if err != nil {
		panic(err.Error())
	}
	index := strings.LastIndex(runPath, string(os.PathSeparator))
	ExEPath = runPath[:index]

	fmt.Println(string(utils.Green), Logo)
	fmt.Println(string(utils.Green), "EXE:"+ExEPath)
}
