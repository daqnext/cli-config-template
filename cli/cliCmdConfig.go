package cli

import (
	"errors"
	fj "github.com/daqnext/fastjson"
	"github.com/daqnext/utils/path_util"
	"github.com/urfave/cli/v2"
	"os"
)

type Cmd struct {
	CmdName    string
	CliContext *cli.Context

	ConfigFile string
	ConfigJson *fj.FastJson
}

var CmdToDo *Cmd

func (c *Cmd) SaveConfigToFile() error {
	return os.WriteFile(c.ConfigFile, c.ConfigJson.GetContent(), 0777)
}

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_LOG = "logs"
const CMD_NAME_SERVICE = "service"

////////config to do cmd ///////////
func configCliCmd() *cli.App {

	var todoerr error

	return &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dev", Required: false},
		},
		Action: func(c *cli.Context) error {
			CmdToDo, todoerr = getCmdToDo(CMD_NAME_DEFAULT, true, c)
			if todoerr != nil {
				return todoerr
			}
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    CMD_NAME_LOG,
				Aliases: []string{CMD_NAME_LOG},
				Usage:   "print all logs ",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "num", Required: false},
					&cli.BoolFlag{Name: "onlyerr", Required: false},
				},
				Action: func(c *cli.Context) error {
					CmdToDo, todoerr = getCmdToDo(CMD_NAME_LOG, false, c)
					if todoerr != nil {
						return todoerr
					}
					return nil
				},
			},
			{
				Name:    CMD_NAME_SERVICE,
				Aliases: []string{CMD_NAME_SERVICE},
				Usage:   "service command",
				Subcommands: []*cli.Command{
					//service install
					{
						Name:  "install",
						Usage: "install meson node in service",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service remove
					{
						Name:  "remove",
						Usage: "remove meson node from service",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service start
					{
						Name:  "start",
						Usage: "run",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service stop
					{
						Name:  "stop",
						Usage: "stop",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service restart
					{
						Name:  "restart",
						Usage: "restart",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service status
					{
						Name:  "status",
						Usage: "show service status",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
				},
			},
		},
	}
}

////////end config to do app ///////////
func readDefaultConfig(isDev bool) (*fj.FastJson, string, error) {
	var defaultConfigPath string
	if isDev {
		LocalLogger.Infoln("======== using dev mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/dev.json")
	} else {
		LocalLogger.Infoln("======== using pro mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/pro.json")
	}

	LocalLogger.Info(defaultConfigPath)

	Config, err := fj.NewFromFile(defaultConfigPath)
	if err != nil {
		LocalLogger.Error("no pro.json under /configs folder , use --dev=true to run dev mode")
		return nil, "", err
	} else {
		return Config, defaultConfigPath, nil
	}
}

func getCmdToDo(cmdName string, needconfig bool, c *cli.Context) (*Cmd, error) {
	path_util.ExEPathPrintln()

	app := &Cmd{
		CmdName:    cmdName,
		CliContext: c,
	}

	if needconfig {
		////read default config
		Config, defaultConfigPath, err := readDefaultConfig(c.Bool("dev"))
		if err != nil {
			return nil, err
		}
		LocalLogger.Infoln("======== start of config ========")
		LocalLogger.Infoln(Config.GetContentAsString())
		LocalLogger.Infoln("======== end of config ========")

		app.ConfigFile = defaultConfigPath
		app.ConfigJson = Config
	}

	logLevel := "INFO"
	if app.ConfigJson != nil {
		var err error
		logLevel, err = app.ConfigJson.GetString("local_log_level")
		if err != nil {
			return nil, errors.New("local_log_level [string] in config not defined," + err.Error())
		}
	}
	err := SetLogLevel(logLevel)
	if err != nil {
		//todo return this error or just log error
		LocalLogger.Errorln(err)
	}
	return app, nil
}

// ManualInitAppConfig init app config when use go test
func ManualInitAppConfig(configPath string) {
	LocalLogger.Infoln("configPath:", configPath)
	Config, err := fj.NewFromFile(configPath)
	if err != nil {
		panic("Manual read config err " + err.Error())
	}

	CmdToDo = &Cmd{
		ConfigFile: configPath,
		ConfigJson: Config,
	}

	logLevel := "INFO"
	logLevel, err = Config.GetString("local_log_level")
	if err != nil {
		panic("local_log_level [string] in config not defined," + err.Error())
	}
	err = SetLogLevel(logLevel)
	if err != nil {
		panic(err.Error())
	}
}
