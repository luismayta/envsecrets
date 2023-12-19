package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	coreconfig "github.com/luismayta/envsecrets/v1/config"
)

func TestConfigLoadEnvSuccess(t *testing.T) {
	conf, err := LoadEnvWithFilename("./mocking/config.env")
	assert.Empty(t, err, err)
	assert.IsType(t, &coreconfig.Config{}, conf)
	assert.Equal(t, "zap", conf.Log.Provider, conf.Log.Provider)
}

func TestConfigMustLoadEnvWithPanic(t *testing.T) {
	assert.Panics(t, func() { MustLoadEnvWithFilename("./mocking/notfound.env") }, "The code did not panic")
}

func TestConfigLoadEnvFailedFailed(t *testing.T) {
	conf, err := LoadEnvWithFilename("./mocking/notfound.env")
	assert.NotEmpty(t, err, err)
	assert.IsType(t, &coreconfig.Config{}, conf)
	assert.Empty(t, conf)
}
