package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
	"path/filepath"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name:    filepath.Base(os.Args[0]),
		Usage:   "The SERO Gate.IO transation analysis tool",
		Version: "0.0.1",
	}

	//app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
