package config

import (
	"github.com/joho/godotenv"

	coreconfig "github.com/luismayta/envsecrets/v1/internal/config"
	"github.com/luismayta/envsecrets/v1/internal/errors"
)

func LoadEnvWithFilename(filename string) (*coreconfig.Config, error) {
	if err := godotenv.Overload(filename); err != nil {
		return nil, errors.Wrapf(err, errors.ErrorNotFound, "filename %s", filename)
	}
	return coreconfig.Must(), nil
}

func MustLoadEnvWithFilename(filename string) *coreconfig.Config {
	conf, err := LoadEnvWithFilename(filename)
	if err != nil {
		panic(err)
	}
	return conf
}
