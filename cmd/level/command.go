package main

import (
	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

var (
	print = &cli.Command{
		Name:   "print",
		Usage:  "prints different log messages by catagory",
		Action: printLog,
	}
)

func printLog(c *cli.Context) error {
	log.Info("just an fyi, Ima 'bout to say something important, so listen up...")
	log.Debugf("bugger off!")
	log.Fatalf("about to die, like everyone and everything")
	log.Errorf("[dead noises...]")
	return nil
}
