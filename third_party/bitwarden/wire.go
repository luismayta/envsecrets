//go:build wireinject

package bitwarden

import (
	"github.com/google/wire"

	"github.com/luismayta/envsecrets/v1/internal/config"
)

func InitializeClient(conf *config.Config) IClient {
	wire.Build(NewBW)
	return nil
}
