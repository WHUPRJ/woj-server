package main

import (
	"github.com/WHUPRJ/woj-server/internal/app"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	a := &cli.App{
		Name:                 "OJ",
		Compiled:             getBuildTime(),
		Version:              Version,
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Paul",
				Email: "i@0x7f.app",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "path to the config file",
				Value:   "config.yaml",
				EnvVars: []string{"APP_CONFIG"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "start the server",
				Action: func(c *cli.Context) error {
					g := new(global.Global)
					g.Setup(c.String("config"))
					defer func() { _ = g.Log.Sync() }()
					//g.SetupRedis()

					g.Log.Info("starting server...")
					return app.Run(g)
				},
			},
		},
	}

	err := a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	BuildTime = "2022-09-06-01-00-00"
	Version   = "1.0.0+None"
)

func getBuildTime() time.Time {
	build, err := time.Parse("2006-01-02-15-04-05", BuildTime)
	if err != nil {
		log.Printf("failed to parse build time: %v", err)
		build = time.Now()
	}
	return build
}
