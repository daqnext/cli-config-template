package cli

import (
	"fmt"

	"github.com/daqnext/cli-config-template/utils"
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

func configCliApp() *cli.App {
	return &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dev", Required: false},
		},
		Action: func(c *cli.Context) error {

			////read default config
			Config, defaultConfigPath, err := readDefaultConfig(c)
			if err != nil {
				return err
			}
			//replace some of the defaultconfig with cli command input
			//flush to defaultconfig.json with overwritten config

			//print config
			fmt.Println(string(utils.Green), "======== using config ========")
			fmt.Println(string(utils.White), Config.GetContentAsString())

			AppToDO = &APP{AppName: "default", ConfigFile: defaultConfigPath, ConfigJson: Config, CliContext: c}
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "logs",
				Aliases: []string{"logs"},
				Usage:   "print all logs ",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "dev", Required: false},
					&cli.IntFlag{Name: "num", Required: false},
					&cli.BoolFlag{Name: "onlyerr", Required: false},
				},
				Action: func(c *cli.Context) error {

					Config, defaultConfigPath, err := readDefaultConfig(c)
					if err != nil {
						return err
					}
					//print config
					fmt.Println(string(utils.Green), "======== using config ========")
					fmt.Println(string(utils.White), Config.GetContentAsString())

					AppToDO = &APP{AppName: "logs", ConfigFile: defaultConfigPath, ConfigJson: Config, CliContext: c}
					return nil
				},
			},
		},
	}
}

func readDefaultConfig(c *cli.Context) (*fj.FastJson, string, error) {
	dev := c.Bool("dev")
	var defaultConfigPath string
	if dev {
		fmt.Println(string(utils.Green), "======== using dev mode ========")
		defaultConfigPath = "config/devconfig.json" // GetPath("config/devconfig.json")
	} else {
		fmt.Println(string(utils.Green), "======== using pro mode ========")
		defaultConfigPath = GetPath("config/proconfig.json")
	}

	Config, err := fj.NewFromFile(defaultConfigPath)
	if err != nil {
		fmt.Println(string(utils.Red), "no proconfig.json , use --dev=true to run dev mode")
		return nil, "", err
	} else {
		return Config, defaultConfigPath, nil
	}
}
