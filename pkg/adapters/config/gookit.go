package config

import (
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type GookitConfig[T interface{}] struct {
	config *T
}

func New[T interface{}](configAddress string) ports.Config[T] {
	if configAddress == "" {
		configAddress = "config.yaml"
	}

	var c *T

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	_ = config.LoadFiles(configAddress)
	config.BindStruct("", c)

	return &GookitConfig[T]{
		config: c,
	}
}

func (gc *GookitConfig[T]) GetConfig() *T {
	return gc.config
}
