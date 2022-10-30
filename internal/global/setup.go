package global

import (
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"log"
)

func (g *Global) SetupZap() {
	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(
			utils.If(g.Conf.Development, zapcore.DebugLevel, zapcore.InfoLevel),
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

func (g *Global) SetupConfig(configFile string) {
	data, err := utils.FileRead(configFile)
	if err != nil {
		log.Fatalf("Failed to setup config: %s\n", err.Error())
	}

	err = yaml.Unmarshal(data, &g.Conf)
	if err != nil {
		log.Fatalf("Failed to setup config: %s\n", err.Error())
	}
}
