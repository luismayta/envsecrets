package config

type Log struct {
	Provider string `env:"LOG_PROVIDER" envDefault:"zap"`
}
