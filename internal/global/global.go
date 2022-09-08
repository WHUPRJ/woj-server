package global

import (
	"github.com/WHUPRJ/woj-server/internal/pkg/metrics"
	"go.uber.org/zap"
)

type Global struct {
	Log  *zap.Logger
	Conf *Config
	Stat *metrics.Metrics
	Db   Repo
}
