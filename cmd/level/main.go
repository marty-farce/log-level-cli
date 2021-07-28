package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

var (
	app = cli.NewApp()
)

func main() {
	app.Usage = "testing log levels"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "log-level",
			Aliases: []string{"l"},
			Usage:   "sets desired log level",
			Value:   "debug",
			EnvVars: []string{"LOG_LEVEL"},
		},
	}
	app.Commands = []*cli.Command{
		print,
	}
	app.Before = func(c *cli.Context) error {
		if logLevel := c.String("log-level"); logLevel != "" {
			if level, err := log.ParseLevel(logLevel); err == nil {
				log.SetLevel(level)
				fmt.Println(level)
			}
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("app.Run ran")
		log.Fatal(err.Error())
	}
}
