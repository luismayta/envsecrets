package log

import (
	"github.com/luismayta/envsecrets/v1/config"
	"github.com/luismayta/envsecrets/v1/internal/common/log/provider"
	"github.com/luismayta/envsecrets/v1/internal/errors"
)

// New initialize a new Log.
func NewLog(conf config.Config) TracingLogger {
	return Factory(conf)
}

// Factory Log.
func Factory(conf config.Config) (prov TracingLogger) {
	switch conf.Log.Provider {
	case "zap":
		prov = provider.NewZap(conf)
	default:
		panic(errors.Errorf(errors.ErrorParseConfig, "unsupported email provider: %s", conf.Log.Provider))
	}
	return prov
}
