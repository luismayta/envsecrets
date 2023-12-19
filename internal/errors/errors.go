package errors

import (
	stderrors "errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Kind is the kind of error.
type Kind string

// Error kinds.
const (
	ErrorReadConfig             Kind = "config read error"
	ErrorParseConfig            Kind = "config parse error"
	ErrorNotImplemented         Kind = "not implement"
	ErrorCanceled               Kind = "canceled"
	ErrorUnknown                Kind = "unknown error"
	ErrorInvalidArgument        Kind = "invalid argument"
	ErrorDeadlineExceeded       Kind = "deadline exceeded"
	ErrorNotFound               Kind = "entity not found"
	ErrorAlreadyExists          Kind = "already exists"
	ErrorPermissionDenied       Kind = "permission denied"
	ErrorUnauthenticated        Kind = "unauthenticated"
	ErrorResourceExhausted      Kind = "too many requests"
	ErrorFailedPrecondition     Kind = "failed precondition"
	ErrorAborted                Kind = "operation aborted"
	ErrorUnavailable            Kind = "service unavailable"
	ErrorUserTemporarilyBlocked Kind = "user has been temporarily blocked"
	ErrorExtensionsDenied       Kind = "extension is denied"
)

// FieldViolation is a struct for providing field error details in HTTP error. It matches the same struct in errdetails package.
type FieldViolation struct {
	Field       string
	Description string
}

// Error is an internal errors with stacktrace. It can be converted to a HTTP response.
type Error struct {
	error
	kind            Kind
	fieldViolations []FieldViolation
}

// Format formats the error.
func (e *Error) Format(s fmt.State, verb rune) {
	var formatter fmt.Formatter
	if errors.As(e.error, &formatter) {
		formatter.Format(s, verb)
		return
	}

	_, err := io.WriteString(s, e.Error())
	if err != nil {
		log.Fatal("Error occurred:", err)
	}
}

// Kind returns the error kind.
func (e *Error) Kind() Kind {
	return e.kind
}

// FieldViolations returns a structure that represents field validation errors.
func (e *Error) FieldViolations() []FieldViolation {
	return e.fieldViolations
}

// New returns an error with the supplied kind and message. If message is empty, a default message,
// for the error kind will be used.
func New(kind Kind, msg string) error {
	if msg == "" {
		msg = string(kind)
	}
	return &Error{
		error: errors.New(msg),
		kind:  kind,
	}
}

// Errorf formats according to a format specifier and return an unknown error with the string.
func Errorf(kind Kind, format string, args ...interface{}) error {
	return New(kind, fmt.Sprintf(format, args...))
}

// Wrap returns an error annotating err with a kind and a stacktrace at the point Wrap is called,
// and the supplied kind and message. If err is nil, Wrap returns nil.
func Wrap(err error, kind Kind, msg string) error {
	if err == nil {
		return nil
	}
	if msg == "" {
		msg = string(kind)
	}
	return &Error{
		error: errors.Wrap(err, msg),
		kind:  kind,
	}
}

// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the
// kind and format specifier. If err is nil, Wrapf returns nil.
func Wrapf(err error, kind Kind, format string, args ...interface{}) error {
	return Wrap(err, kind, fmt.Sprintf(format, args...))
}

// Must returns an error annotating err with a kind and a stacktrace at the point Must is called,
func Must(err error, kind Kind, format string, args ...interface{}) {
	err = Wrap(err, kind, fmt.Sprintf(format, args...))
	if err != nil {
		panic(err)
	}
}

// WithFieldViolations returns an error with supplied field
// violations.
func WithFieldViolations(kind Kind, msg string, fieldViolations []FieldViolation) error {
	if msg == "" {
		msg = string(kind)
	}
	return &Error{
		error:           errors.New(msg),
		kind:            kind,
		fieldViolations: fieldViolations,
	}
}

// WithValidateError maps a Validate error into an internal error representation.
func WithValidateError(err error) error {
	if err == nil {
		return nil
	}

	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {
		var fieldViolations []FieldViolation
		for _, fieldError := range fieldErrors {
			fieldViolations = append(fieldViolations, FieldViolation{
				Field:       ToUnderScore(fieldError.Field()),
				Description: fieldError.Tag(),
			})
		}
		return WithFieldViolations(ErrorInvalidArgument, "", fieldViolations)
	}

	return Wrap(err, ErrorUnknown, "")
}

// HTTPStatusCodeFromKind converts an error kind into HTTP status code.
func HTTPStatusCodeFromKind(kind Kind) int {
	statusCodeMap := map[Kind]int{
		ErrorCanceled:           http.StatusRequestTimeout,
		ErrorUnknown:            http.StatusInternalServerError,
		ErrorInvalidArgument:    http.StatusBadRequest,
		ErrorDeadlineExceeded:   http.StatusGatewayTimeout,
		ErrorNotFound:           http.StatusNotFound,
		ErrorAlreadyExists:      http.StatusConflict,
		ErrorPermissionDenied:   http.StatusForbidden,
		ErrorUnauthenticated:    http.StatusUnauthorized,
		ErrorResourceExhausted:  http.StatusTooManyRequests,
		ErrorFailedPrecondition: http.StatusBadRequest,
		ErrorAborted:            http.StatusConflict,
		ErrorUnavailable:        http.StatusServiceUnavailable,
	}

	if code, ok := statusCodeMap[kind]; ok {
		return code
	}

	log.Infof("Unknown error kind: %v", kind)
	return http.StatusInternalServerError
}

// IsKind checks whether any error in err's chain matches the error kind.
func IsKind(err error, kind Kind) bool {
	ie := &Error{}
	if As(err, &ie) {
		return ie.kind == kind
	}
	return false
}

// As finds the first error in err's chain that matches target, and if so, sets target to that
// error value and return true.
func As(err error, target interface{}) bool {
	return stderrors.As(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// Same as Go's errors.Unwrap.
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}
