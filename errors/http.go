package errors

import (
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

// 类型扩展方法:
type HttpError errors.Error

func (m *HttpError) Error() string {
	// 强制类型转换 + 调用
	return (*errors.Error)(m).Error()
}

func (m *HttpError) GetCode() int32 {
	// 强制类型转换 + 调用
	return (*errors.Error)(m).GetCode()
}

func (m *HttpError) GetDetail() string {
	return (*errors.Error)(m).GetDetail()
}

func (m *HttpError) GetId() string {
	return (*errors.Error)(m).GetId()
}

// New generates a custom error.
func New(id, detail string, code int32) error {
	return &errors.Error{
		Id:     id,
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// BadRequest generates a 400 error.
func BadRequest(id, format string, a ...interface{}) error {
	return errors.BadRequest(id, format, a...)
}

// Unauthorized generates a 401 error.
func Unauthorized(id, format string, a ...interface{}) error {
	return errors.Unauthorized(id, format, a...)
}

// Forbidden generates a 403 error.
func Forbidden(id, format string, a ...interface{}) error {
	return errors.Forbidden(id, format, a...)
}

// NotFound generates a 404 error.
func NotFound(id, format string, a ...interface{}) error {
	return errors.NotFound(id, format, a...)
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(id, format string, a ...interface{}) error {
	return errors.MethodNotAllowed(id, format, a...)
}

// Timeout generates a 408 error.
func Timeout(id, format string, a ...interface{}) error {
	return errors.Timeout(id, format, a...)
}

// Conflict generates a 409 error.
func Conflict(id, format string, a ...interface{}) error {
	return errors.Conflict(id, format, a...)
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, a ...interface{}) error {
	return errors.InternalServerError(id, format, a...)
}
