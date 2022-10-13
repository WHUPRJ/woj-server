package main

import (
	"github.com/WHUPRJ/woj-server/global"
	"github.com/WHUPRJ/woj-server/internal/app"
	"github.com/urfave/cli/v2"
	"math/rand"
	"time"
)

func run(c *cli.Context) error {
	rand.Seed(time.Now().Unix())

	g := new(global.Global)
	g.SetupConfig(c.String("config"))
	g.SetupZap()
	defer func() { _ = g.Log.Sync() }()

	g.Log.Info("starting server...")
	return app.Run(g)
}
