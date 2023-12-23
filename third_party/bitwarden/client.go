package bitwarden

import (
	"github.com/luismayta/envsecrets/v1/internal/config"
)

func NewClient(conf *config.Config) IClient {
	return NewBW(conf)
}
