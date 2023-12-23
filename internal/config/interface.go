package config

// Configurer methods for config.
type Configurer interface {
	ReadConfig() (*Config, error)
}
