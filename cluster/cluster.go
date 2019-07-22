package cluster

import (
	"github.com/2d0g/influxdb-srelay/config"
)

var (
	mainConfig *config.Config
	logDir     string
)

func SetConfig(cfg *config.Config) {
	mainConfig = cfg
}

func SetLogdir(ld string) {
	logDir = ld
}
