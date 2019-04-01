package errors

import (
	"fmt"
	"net/http"
)

// Error wraps up a message and an http status code in ordre to pass
// the relevant info up to the top level controller
// This class is really just a bunch of convenience methods
type Error struct {
	msg  string
	code int
}

func (e Error) Error() string {
	return e.msg
}

func (e Error) String() string {
	return e.msg
}

func (e Error) Code() int {
	return e.code
}

func New(code int, msg string) Error {
	return Error{
		msg:  msg,
		code: code,
	}
}

func Newf(code int, msg string, args ...interface{}) Error {
	return Error{
		code: code,
		msg:  fmt.Sprintf(msg, args...),
	}
}

func Unauthorized(msg string) Error {
	return New(http.StatusUnauthorized, msg)
}

func Unauthorizedf(msg string, args ...interface{}) Error {
	return Newf(http.StatusUnauthorized, msg, args...)
}

func BadRequest(msg string) Error {
	return New(http.StatusBadRequest, msg)
}

func BadRequestf(msg string, args ...interface{}) Error {
	return Newf(http.StatusBadRequest, msg, args...)
}

func NotFound(msg string) Error {
	return New(http.StatusNotFound, msg)
}

func NotFoundf(msg string, args ...interface{}) Error {
	return Newf(http.StatusNotFound, msg, args...)
}

func Internal(msg string) Error {
	return New(http.StatusInternalServerError, msg)
}

func Internalf(msg string, args ...interface{}) Error {
	return Newf(http.StatusInternalServerError, msg, args...)
}
