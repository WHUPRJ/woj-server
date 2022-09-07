package global

import (
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func (g *Global) setupZap() {
	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(
			utils.If(g.Conf.Development, zapcore.DebugLevel, zapcore.InfoLevel).(zapcore.Level),
		),
		Development:      g.Conf.Development,
		Encoding:         "console", // or json
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	g.Log, err = cfg.Build()
	if err != nil {
		log.Fatalf("Failed to setup Zap: %s\n", err.Error())
	}
}
