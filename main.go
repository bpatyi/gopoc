package main

import (
	"flag"
	"fmt"

	"github.com/bpatyi/gopoc/app"
)

var configPath = flag.String("config", "config.json", "Path to configuration file")
var configEnvPrefix = flag.String("config_env_prefix", "gopoc", "Environment prefix of configs")

func main() {
	flag.Parse()

	application, err := app.NewApp(configPath, configEnvPrefix)
	if err != nil {
		panic(fmt.Errorf("Failed to create application: %w", err))
	}
	_ = application
}
