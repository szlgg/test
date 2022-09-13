package gerroraa

import "fmt"

// Newf returns an error that formats as the given format and args.
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  -1,
	}
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// It returns nil if given <err> is nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// New creates and returns an error which is formatted from given text.
func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  -1,
	}
}

// NewSkipf returns an error that formats as the given format and args.
// The parameter <skip> specifies the stack callers skipped amount.
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  -1,
	}
}

// Error implements the interface of Error, it returns all the error as string.
func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if err.error != nil {
		if err.text != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}

// Code returns the error code.
// It returns -1 if it has no error code.
func (err *Error) Code() int {
	if err == nil {
		return -1
	}
	return err.code
}

// Code returns the error code of current error.
// It returns -1 if it has no error code or it does not implements interface Code.
func Code(err error) int {
	if err != nil {
		if e, ok := err.(apiCode); ok {
			return e.Code()
		}
	}
	return -1
}

// apiCode is the interface for Code feature.
type apiCode interface {
	Error() string // It should be an error.
	Code() int
}