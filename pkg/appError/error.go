package appError

import "net/http"

var (
	ErrorNotFound  = NotFound("not found")
	ErrorInternal  = InternalError("internal error")
	ErrorDuplicate = BadRequest("duplicate")
)

type Error struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NotFound(message string) *Error {
	return New(http.StatusNotFound, message)
}

func BadRequest(message string) *Error {
	return New(http.StatusBadRequest, message)
}

func InternalError(message string) *Error {
	return New(http.StatusInternalServerError, message)
}
