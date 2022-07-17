package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Load(configPath, configEnvPrefix *string) (*Config, error) {
	var config Config

	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configEnvPrefix != nil {
		v.SetEnvPrefix(*configEnvPrefix)
	}
	if configPath != nil {
		v.SetConfigFile(*configPath)

		err := v.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("the given config file is not found: %w", err)
		}
	}

	v.Unmarshal(&config)
	v.AutomaticEnv()

	return &config, config.Validate()
}
