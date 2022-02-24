package errors

import "github.com/yuonoda/bookspace/app/errors/codes"

type Error struct {
	Code codes.Code
}

func (e Error) Error() string {
	return ""
}

func Invalid() Error {
	return Error{
		Code: codes.Invalid,
	}
}

func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if e, ok := err.(Error); ok {
		return e.Code
	}
	return codes.InternalError
}
