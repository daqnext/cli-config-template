package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func readDefaultConfig(c *cli.Context) (*fj.FastJson, string, error) {
	dev := c.Bool("dev")
	var defaultConfigPath string
	if dev {
		fmt.Println(string(utils.Green), "======== using dev mode ========")
		defaultConfigPath = GetPath("config/devconfig.json")
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

var ExEPath string

func GetPath(relpath string) string {
	return ExEPath + "/" + strings.Trim(relpath, "/")
}

func init() {

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

	//print any initialzation panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(utils.Red), "panic errors:", err.(error).Error())
		}
	}()

	CliApp := &cli.App{
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

	errRun := CliApp.Run(os.Args)
	if errRun != nil {
		log.Fatal(errRun)
		panic(errRun.Error())
	}

}
