package config

import (
	"github.com/abtinokhovat/gox/confx"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

type ComplexConfig struct {
	SomeOtherFields string `koanf:"some_other_fields" secret:"true"`
}

type Config struct {
	SomeConfig    string        `koanf:"some_config"`
	SomeSecret    string        `koanf:"some_secret" secret:"true"`
	ComplexConfig ComplexConfig `koanf:"complex_config"`
}

func (c Config) String() string {
	return confx.MaskSecrets(c)
}

func Default() Config {
	return Config{
		SomeConfig: "some",
		SomeSecret: "2983479hfsiehfiash9238473",
		ComplexConfig: ComplexConfig{
			SomeOtherFields: "some",
		},
	}
}

// Get retrieves the singleton instance of the config.
func Get() (*Config, error) {
	var loadErr error

	// load config from and return it
	once.Do(func() {
		opts := confx.Option{
			Prefix:    "ABTIN_",
			Delimiter: ".",
			Separator: "__",
		}

		instance = new(Config)

		loadErr = confx.NewConfigLoader[Config](opts).WithDefaultProvider(Default()).WithEnvProvider().WithYamlProvider("").Load(instance)
	})

	if loadErr != nil {
		return nil, loadErr
	}

	return instance, nil
}
