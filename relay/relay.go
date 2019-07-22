package relay

import (
	"fmt"
	"github.com/2d0g/influxdb-srelay/cluster"
	"github.com/2d0g/influxdb-srelay/config"
)

// Relay is an HTTP or UDP endpoint
/*type Relay interface {
	Name() string
	Run() error
	Stop() error
}*/

var (
	mainConfig *config.Config
	logDir     string
	clusters   map[string]*cluster.Cluster
)

func SetConfig(cfg *config.Config) {
	mainConfig = cfg
}

func SetLogdir(ld string) {
	logDir = ld
}

func InitClusters() error {

	clusters = make(map[string]*cluster.Cluster)

	for _, cfg := range mainConfig.Influxcluster {

		c, err := cluster.NewCluster(cfg)
		if err != nil {
			return err
		}
		if clusters[cfg.Name] != nil {
			return fmt.Errorf("duplicate cluster: %q", cfg.Name)
		}
		clusters[cfg.Name] = c
	}
	return nil
}
