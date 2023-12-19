package log

// import (
// 	"testing"

// 	"github.com/luismayta/envsecrets/v1/internal/testutil/config"
// 	"github.com/stretchr/testify/assert"
// )

// func logForTest() (TracingLogger, func()) {
// 	conf := config.MustLoadEnvWithFilename("./mocking/zap.env")

// 	log := NewLog(*conf)

// 	return log, func() {}
// }

// func TestNewSuccess(t *testing.T) {
// 	log, tearDown := logForTest()
// 	defer tearDown()
// 	assert.IsType(t, &TracingLogger{}, log)
// }
