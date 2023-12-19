package provider

import (
	"go.uber.org/zap"

	"github.com/luismayta/envsecrets/config"
)

// Zap is a struct of Zap.
type ZapLog struct {
	Client *zap.Logger
}

func NewZap(conf config.Config) *ZapLog {
	logger, _ := zap.NewProduction()
	defer func() { // flushes buffer, if any
		_ = logger.Sync()
	}()
	return &ZapLog{Client: logger}
}

func (l *ZapLog) Close() {
	defer func() { // flushes buffer, if any
		_ = l.Client.Sync()
	}()
}

// Send an debug.
func (l *ZapLog) Debugf(msg string, args ...interface{}) {
	l.Client.Sugar().Debugf(msg, args)
	defer func() { // flushes buffer, if any
		_ = l.Client.Sync()
	}()
}

// Send an infof.
func (l *ZapLog) Infof(msg string, args ...interface{}) {
	l.Client.Sugar().Infof(msg, args)
	defer func() { // flushes buffer, if any
		_ = l.Client.Sync()
	}()
}

// Send an error.
func (l *ZapLog) Error(msg string) {
	l.Client.Sugar().Error(msg)
	defer func() { // flushes buffer, if any
		_ = l.Client.Sync()
	}()
}
