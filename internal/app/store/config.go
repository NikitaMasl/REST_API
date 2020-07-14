package store

// Config ...
type Config struct {
	Database string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
