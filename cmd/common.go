package cmd

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
	"time"
)

var App = &cli.App{
	Name:                 "WOJ",
	EnableBashCompletion: true,
	Authors: []*cli.Author{
		{
			Name:  "Paul",
			Email: "i@0x7f.app",
		},
		{
			Name:  "cxy004",
			Email: "cxy004@qq.com",
		},
		{
			Name:  "wzt",
			Email: "w.zhongtao@qq.com",
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
}

var (
	BuildTime string
	Version   string
)

func init() {
	if BuildTime == "" {
		BuildTime = "2022-09-06-01-00-00"
	}
	App.Compiled = getBuildTime()

	if Version == "" {
		Version = "0.0.0+None"
	}
	App.Version = Version
}

func getBuildTime() time.Time {
	build, err := time.Parse("2006-01-02-15-04-05", BuildTime)
	if err != nil {
		log.Printf("failed to parse build time: %v", err)
		build = time.Now()
	}
	return build
}

func CommonSetup(c *cli.Context) *global.Global {
	rand.Seed(time.Now().Unix())

	g := new(global.Global)
	g.SetupConfig(c.String("config"))
	g.SetupZap()

	g.Log.Info("starting...")

	return g
}
