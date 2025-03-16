## Config Loader

The `config` package is responsible for loading configuration values from multiple sources, including environment variables, YAML files, and default values. The configuration values are managed using the `koanf` library, allowing flexible configuration loading and merging.

### API

**`C() *Config`** → A singleton function to get the loaded config through the application wherever needed.

**`NewConfigLoader(opts Option) *Loader`** → Creates a new `Loader` instance with the specified options.

**`WithEnvProvider(callback *EnvCallbackFunc) *Loader`** → Adds an environment variable provider to the loader.
- **callback**: A function that processes environment variable names. If not provided, a default callback function will be used.

**`WithYamlProvider(filePath string) *Loader`** → Adds a YAML file provider to the loader.
- **filePath**: Path to the YAML file from which the configuration will be loaded. If not provided, it defaults to `./config.yml`.

**`WithDefaultProvider() *Loader`** →  Adds the default provider for structs to the loader. This provides default values for configuration if no other sources are loaded.

**`Load() (*Config, error)`** →  Loads the configuration into a provided `Config` struct, unmarshaling the loaded values into it.

### Default Environment Variable Format
The environment variable names follow the format `ORDER_<SECTION>__<KEY>`, where:

- `ORDER_`: Default prefix for environment variables.
- `__`: Separator used for nesting configuration fields (e.g., `ORDER_SIMULATE__LOG_LEVEL`).
- `.`: Delimiter used in the configuration structure.

The loader converts the environment variable names into appropriate keys based on the separator and delimiter.
