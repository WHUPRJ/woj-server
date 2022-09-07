package global

import (
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"gopkg.in/yaml.v3"
	"log"
)

type ConfigWebServer struct {
	Address string `yaml:"Address"`
	Port    int    `yaml:"Port"`
}

type ConfigRedis struct {
	Db       int    `yaml:"Db"`
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
}

type ConfigMetrics struct {
	Namespace string `yaml:"Namespace"`
	Subsystem string `yaml:"Subsystem"`
}

type Config struct {
	WebServer   ConfigWebServer `yaml:"WebServer"`
	Redis       ConfigRedis     `yaml:"Redis"`
	Metrics     ConfigMetrics   `yaml:"Metrics"`
	Development bool            `yaml:"Development"`
}

func (g *Global) setupConfig(configFile string) {
	data, err := utils.FileRead(configFile)
	if err != nil {
		log.Fatalf("Failed to setup config: %s\n", err.Error())
	}

	err = yaml.Unmarshal(data, &g.Conf)
	if err != nil {
		log.Fatalf("Failed to setup config: %s\n", err.Error())
	}
}
