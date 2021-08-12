package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
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

	gflags := []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "optionbool"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "optionnum"}),
	}

	app := &cli.App{

		Action: func(c *cli.Context) error {
			fmt.Println("this is default command action")
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "firstcmd",
				Aliases: []string{"fc"},
				Usage:   "first command ",
				Flags:   gflags,
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
				Before: altsrc.InitInputSource(gflags, func() (altsrc.InputSourceContext, error) {
					return altsrc.NewJSONSourceFromFile(fileurl)
				}),
			},

			{
				Name:    "second",
				Aliases: []string{"sec"},
				Usage:   "second command ",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "param1"},
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
