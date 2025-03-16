package appa

import (
	"github.com/abtinokhovat/gox/config"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	SomeConfig string `koanf:"some_config"`
}

// Get retrieves the singleton instance of the loader.
func Get() (*Config, error) {
	var loadErr error

	// load config from and return it
	once.Do(func() {
		opts := config.Option{
			Prefix:    "CORE",
			Delimiter: ".",
			Separator: "__",
		}
		loadErr = config.NewConfigLoader[Config](opts).WithDefaultProvider(Default()).Load(instance)
	})

	if loadErr != nil {
		return nil, loadErr
	}

	return instance, nil
}
