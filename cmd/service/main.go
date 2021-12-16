package service

import (
	"fmt"
	clitool "github.com/daqnext/cli-config-template/cli"
	"github.com/daqnext/daemon"
	"log"
	"os"
	"runtime"
)

const (
	// name of the service
	name        = "cli-config-template"
	description = "cli-config-template description"
)

type Service struct {
	daemon.Daemon
}

func RunServiceCmd() {
	//check command
	subCmds := clitool.CmdToDo.CliContext.Command.Names()
	if len(subCmds) == 0 {
		clitool.LocalLogger.Fatalln("no sub command")
		return
	}

	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		clitool.LocalLogger.Fatalln("run daemon error:", err)
	}
	service := &Service{srv}

	action := subCmds[0]
	log.Println(action)

	var status string
	var e error
	switch action {
	case "install":
		status, e = service.Install()
		clitool.LocalLogger.Debugln("cmd install")
	case "remove":
		service.Stop()
		status, e = service.Remove()
		clitool.LocalLogger.Debugln("cmd remove")
	case "start":
		status, e = service.Start()
		clitool.LocalLogger.Debugln("cmd start")
	case "stop":
		status, e = service.Stop()
		clitool.LocalLogger.Debugln("cmd stop")
	case "restart":
		service.Stop()
		status, e = service.Start()
		clitool.LocalLogger.Debugln("cmd restart")
	case "status":
		status, e = service.Status()
		clitool.LocalLogger.Debugln("cmd status")
	default:
		clitool.LocalLogger.Debugln("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}
