package main

import (
	"github.com/WHUPRJ/woj-server/cmd"
	"github.com/WHUPRJ/woj-server/internal/app/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	a := cmd.App
	a.Usage = "woj-server"
	a.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "start the server",
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

	return server.RunServer(g)
}
