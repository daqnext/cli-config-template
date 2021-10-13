package cli

import (
	fj "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

type APP struct {
	ConfigFile string
	AppName    string
	ConfigJson *fj.FastJson
	CliContext *cli.Context
}

var AppToDO *APP

const APP_NAME_DEFAULT = "default"
const APP_NAME_LOG = "logs"

func AppIsActive(appName string) bool {
	if AppToDO.AppName == appName {
		return true
	} else {
		return false
	}
}

////////config to do app ///////////
func configCliApp() *cli.App {

	var todoerr error

	return &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dev", Required: false},
		},
		Action: func(c *cli.Context) error {
			AppToDO, todoerr = getAppToDo(APP_NAME_DEFAULT, true, c)
			if todoerr != nil {
				return todoerr
			}
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    APP_NAME_LOG,
				Aliases: []string{APP_NAME_LOG},
				Usage:   "print all logs ",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "num", Required: false},
					&cli.BoolFlag{Name: "onlyerr", Required: false},
				},
				Action: func(c *cli.Context) error {
					AppToDO, todoerr = getAppToDo(APP_NAME_LOG, false, c)
					if todoerr != nil {
						return todoerr
					}
					return nil
				},
			},
		},
	}
}

////////end config to do app ///////////

func readDefaultConfig(configPrefixPath string, c *cli.Context) (*fj.FastJson, string, error) {
	dev := c.Bool("dev")
	var defaultConfigPath string
	if dev {
		LocalLogger.Infoln("======== using dev mode ========")
		defaultConfigPath = configPrefixPath + "devconfig.json"
	} else {
		LocalLogger.Infoln("======== using pro mode ========")
		defaultConfigPath = configPrefixPath + "proconfig.json"
	}

	LocalLogger.Info(defaultConfigPath)

	Config, err := fj.NewFromFile(defaultConfigPath)
	if err != nil {
		LocalLogger.Error("no proconfig.json , use --dev=true to run dev mode")
		return nil, "", err
	} else {
		return Config, defaultConfigPath, nil
	}
}

func getAppToDo(appName string, needconfig bool, c *cli.Context) (*APP, error) {
	if needconfig {
		LocalLogger.Infoln("EXE:" + ExEPath)
		////read default config
		Config, defaultConfigPath, err := readDefaultConfig(GetPath("configs/"+appName+"_"), c)
		if err != nil {
			return nil, err
		}
		LocalLogger.Infoln("======== start of config ========")
		LocalLogger.Infoln(Config.GetContentAsString())
		LocalLogger.Infoln("======== end of config ========")
		return &APP{AppName: appName, ConfigFile: defaultConfigPath, ConfigJson: Config, CliContext: c}, nil
	} else {
		return &APP{AppName: appName, ConfigFile: "", ConfigJson: nil, CliContext: c}, nil
	}

}
