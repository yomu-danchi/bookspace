package errors

import (
	"github.com/yuonoda/bookspace/app/errors/codes"
)

type AppError struct {
	Code  codes.Code
	error error
}

func (e AppError) Error() string {
	return ""
}

func Invalid(err error) AppError {
	return AppError{
		error: err,
		Code:  codes.Invalid,
	}
}

func NotFound(err error) AppError {
	return AppError{
		error: err,
		Code:  codes.NotFound,
	}
}

func Internal(err error) AppError {
	return AppError{
		error: err,
		Code:  codes.InternalError,
	}
}

func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if e, ok := err.(AppError); ok {
		return e.Code
	}
	return codes.InternalError
}
