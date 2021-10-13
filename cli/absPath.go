package cli

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
}
