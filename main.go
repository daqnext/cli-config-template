package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/buger/jsonparser"
	"github.com/urfave/cli/v2"
)

func getGlobalJson(fileurl string) (*gabs.Container, error) {

	byte, err := ioutil.ReadFile(fileurl) // just pass the file name
	if err != nil {
		return nil, err
	}
	jsonParsed, errjson := gabs.ParseJSON(byte)

	if errjson != nil {
		return nil, errjson
	}
	return jsonParsed, nil
}

func setGlobalJson(json *gabs.Container, fileurl string) {
	configFile, err := os.OpenFile(fileurl, os.O_RDWR, 0666)
	defer configFile.Close()
	if err != nil {
		return
	}
	configFile.Truncate(0)
	configFile.Seek(0, 0)
	configFile.Write(json.Bytes())
	configFile.Sync()
}

func main() {

	fileurl := "gconfig.json"

	app := &cli.App{

		Action: func(c *cli.Context) error {
			fmt.Println("this is default command action")

			jdata, _err := ioutil.ReadFile(fileurl)
			if _err == nil {
				optbool, _ := jsonparser.GetBoolean(jdata, "optionbool")
				fmt.Println("optionbool:", optbool)
				optnum, _ := jsonparser.GetInt(jdata, "optionnum")
				fmt.Println("optnum:", optnum)
			}
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

					dyj, _err := getGlobalJson(fileurl)
					if _err == nil {
						dyj.Set(c.Bool("optionbool"), "optionbool")
						dyj.Set(c.Int("optionnum"), "optionnum")
						setGlobalJson(dyj, fileurl)
					}
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
