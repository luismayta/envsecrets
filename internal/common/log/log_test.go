package log

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/luismayta/envsecrets/v1/internal/testutil/config"
)

func logForTest() (TracingLogger, func()) {
	conf := config.MustLoadEnvWithFilename("./mocking/zap.env")

	log := NewLog(*conf)

	return log, func() {}
}

func TestNewSuccess(t *testing.T) {
	log, tearDown := logForTest()
	defer tearDown()
	assert.NotNil(t, log, "log should not be nil")
}
