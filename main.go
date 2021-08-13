package main

import (
	"fmt"
	"log"
	"os"

	FastJson "github.com/daqnext/fastjson"
	"github.com/urfave/cli/v2"
)

// func overwriteConfig(fileurl string, content []byte) {
// 	configFile, err := os.OpenFile(fileurl, os.O_RDWR, 0666)
// 	defer configFile.Close()
// 	if err != nil {
// 		return
// 	}
// 	configFile.Truncate(0)
// 	configFile.Seek(0, 0)
// 	configFile.Write(content)
// 	configFile.Sync()
// }

func main() {

	fileurl := "gconfig.json"

	app := &cli.App{

		Action: func(c *cli.Context) error {
			fmt.Println("this is default command action")

			fj, err := FastJson.NewFromFile(fileurl)
			if err != nil {
				return err
			}

			op_bool, _ := fj.GetBoolean("optionbool")
			op_num, _ := fj.GetInt("optionnum")

			fmt.Println("optionbool:", op_bool)
			fmt.Println("optionnum:", op_num)

			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "firstcmd",
				Aliases: []string{"fc"},
				Usage:   "first command ",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "optionbool", Required: true},
					&cli.IntFlag{Name: "optionnum", Required: true},
				},
				Action: func(c *cli.Context) error {

					fmt.Println("optionbool:", c.Bool("optionbool"))
					fmt.Println("optionnum:", c.Int("optionnum"))

					fj, err := FastJson.NewFromFile(fileurl)
					if err != nil {
						return err
					}
					fj.SetBoolean(c.Bool("optionbool"), "optionbool")
					fj.SetInt(c.Int("optionnum"), "optionnum")
					fj.ClearFileAndOutput(fileurl)

					return nil
				},
			},

			{
				Name:    "second",
				Aliases: []string{"sec"},
				Usage:   "second command ",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "param1", Required: true},
					&cli.StringFlag{Name: "param2"},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("param1:", c.String("param1"))
					fmt.Println("param2:", c.String("param2"))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
