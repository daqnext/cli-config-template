package cli

import (
	"fmt"
	"log"
	"os"

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

func init() {

	fmt.Println(string(utils.Green), Logo)

	CliApp := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dev", Required: false},
		},
		Action: func(c *cli.Context) error {

			////read default config
			dev := c.Bool("dev")
			var defaultConfigPath string
			if dev {
				defaultConfigPath = "config/devconfig.json"
			} else {
				defaultConfigPath = "config/proconfig.json"
			}

			Config, err := fj.NewFromFile(defaultConfigPath)
			if err != nil {
				return err
			}

			//replace some of the globalconfig with cli command input

			//flush to globalconfig.json with overwritten config

			//print config
			fmt.Println(string(utils.Green), "======== using config ========")
			fmt.Println(string(utils.Purple), Config.GetContentAsString())

			AppToDO = &APP{AppName: "default", ConfigFile: defaultConfigPath, ConfigJson: Config, CliContext: c}
			return nil
		},

		// Commands: []*cli.Command{
		// 	{
		// 		Name:    "firstcmd",
		// 		Aliases: []string{"fc"},
		// 		Usage:   "first command ",
		// 		Flags: []cli.Flag{
		// 			&cli.BoolFlag{Name: "optionbool", Required: true},
		// 			&cli.IntFlag{Name: "optionnum", Required: true},
		// 		},
		// 		Action: func(c *cli.Context) error {

		// 			fmt.Println("optionbool:", c.Bool("optionbool"))
		// 			fmt.Println("optionnum:", c.Int("optionnum"))

		// 			fj, err := FastJson.NewFromFile(fileurl)
		// 			if err != nil {
		// 				return err
		// 			}
		// 			fj.SetBoolean(c.Bool("optionbool"), "optionbool")
		// 			fj.SetInt(c.Int64("optionnum"), "optionnum")
		// 			fj.ClearFileAndOutput(fileurl)

		// 			return nil
		// 		},
		// 	},

		// 	{
		// 		Name:    "second",
		// 		Aliases: []string{"sec"},
		// 		Usage:   "second command ",
		// 		Flags: []cli.Flag{
		// 			&cli.StringFlag{Name: "param1", Required: true},
		// 			&cli.StringFlag{Name: "param2"},
		// 		},
		// 		Action: func(c *cli.Context) error {
		// 			fmt.Println("param1:", c.String("param1"))
		// 			fmt.Println("param2:", c.String("param2"))
		// 			return nil
		// 		},
		// 	},
		// },
	}

	err := CliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
