package provider

import (
	"testing"

	"github.com/luismayta/envsecrets/internal/testutil/config"
)

func zapForTest() (*ZapLog, func()) {
	conf := config.MustLoadEnvWithFilename("./mocking/zap.env")
	log := NewZap(*conf)

	return log, func() {}
}

func TestZapDebugSuccess(t *testing.T) {
	log, tearDown := zapForTest()
	defer tearDown()

	log.Debugf("test subject")
}
