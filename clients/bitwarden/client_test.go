package bitwarden

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/luismayta/envsecrets/v1/internal/app/config"
)

func TestNewClient(t *testing.T) {
	conf := config.Initialize()
	client := NewClient(conf)
	assert.NotNil(t, client)

	_, ok := client.(IClient)
	assert.True(t, ok, "log should implement Client interface")
}
