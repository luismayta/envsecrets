package bitwarden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)

	_, ok := client.(Client)
	assert.True(t, ok, "log should implement Client interface")
}