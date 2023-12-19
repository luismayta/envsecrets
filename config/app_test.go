package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppVersionNotNil(t *testing.T) {
	conf := Initialize()
	assert.NotEmpty(t, conf.App.Version)
	assert.NotEmpty(t, conf.App.Name, conf.App.Name)
	assert.NotEmpty(t, conf.App.Description, conf.App.Description)
	assert.NotEmpty(t, conf.App.Version, conf.App.Version)
}
