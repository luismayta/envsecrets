package errors

import (
	stderrors "errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	err := New(ErrorNotFound, "internal error")
	assert.Equal(t, "internal error", fmt.Sprintf("%v", err))
	assert.Contains(t, fmt.Sprintf("%+v", err), "internal error", err) // with stacktrace
}

func TestFormatEmpty(t *testing.T) {
	err := New(ErrorPermissionDenied, "")
	assert.Equal(t, "permission denied", fmt.Sprintf("%v", err))
	assert.Contains(t, fmt.Sprintf("%+v", err), "permission denied", err) // with stacktrace
}

func TestNotReadConfig(t *testing.T) {
	err := New(ErrorReadConfig, "")
	assert.Equal(t, "config read error", fmt.Sprintf("%v", err))
	assert.Contains(t, fmt.Sprintf("%+v", err), "config read error", err) // with stacktrace
}

func TestNotParseConfig(t *testing.T) {
	err := New(ErrorParseConfig, "")
	assert.Equal(t, "config parse error", fmt.Sprintf("%v", err))
	assert.Contains(t, fmt.Sprintf("%+v", err), "config parse error", err) // with stacktrace
}

func TestIsKind(t *testing.T) {
	err := New(ErrorNotFound, "internal error")
	assert.True(t, IsKind(err, ErrorNotFound))

	err = stderrors.New("std error")
	assert.False(t, IsKind(err, ErrorNotFound))

	err = nil
	assert.False(t, IsKind(err, ErrorNotFound))
}

func TestAs(t *testing.T) {
	err := New(ErrorNotFound, "not found")
	err2 := &Error{}
	assert.True(t, As(err, &err2))
	assert.Equal(t, ErrorNotFound, err2.kind)
}
