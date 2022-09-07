package global

import (
	"github.com/WHUPRJ/woj-server/internal/pkg/metrics"
	"math/rand"
	"time"
)

func (g *Global) Setup(configFile string) {
	rand.Seed(time.Now().Unix())

	g.setupConfig(configFile)
	g.setupZap()

	g.Stat = new(metrics.Metrics)
	g.Stat.Setup(g.Conf.Metrics.Namespace, g.Conf.Metrics.Subsystem)
	g.Stat.SetLogPaths([]string{
		"/api",
	})
}
