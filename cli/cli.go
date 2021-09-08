package cli

import (
	"fmt"
	"log"
	"os"

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

	//fmt.Println("cli.go init")

	CliApp := &cli.App{
		Action: func(c *cli.Context) error {

			////read default config
			defaultConfigPath := "config/globalconfig.json"
			Config, err := fj.NewFromFile(defaultConfigPath)
			if err != nil {
				return err
			}

			//replace some of the globalconfig with cli command input

			//flush to config.json

			//print config
			fmt.Println("======== using config ========")
			fmt.Println(Config.GetContentAsString())

			AppToDO = &APP{ConfigFile: defaultConfigPath, AppName: "default", ConfigJson: Config, CliContext: c}
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
