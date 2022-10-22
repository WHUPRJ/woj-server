package main

import (
	"github.com/WHUPRJ/woj-server/cmd"
	"github.com/WHUPRJ/woj-server/internal/app/runner"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	a := cmd.App
	a.Usage = "woj-runner"
	a.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "start the runner",
			Action:  run,
		},
	}

	err := a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	g := cmd.CommonSetup(c)
	defer func() { _ = g.Log.Sync() }()

	return runner.RunRunner(g)
}
