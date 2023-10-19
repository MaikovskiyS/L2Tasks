package apperror

import (
	"encoding/json"
	"errors"
)

type errorCode string

var (
	ErrNotFound   errorCode = "notFound"
	ErrInternal   errorCode = "internalErr"
	ErrBadRequest errorCode = "badRequest"
)
var ErrMethod error = errors.New("method incorrect")

type AppError struct {
	Code     errorCode
	Location string
	err      error
}

func (er *AppError) Error() string {
	return er.err.Error()
}
func (er *AppError) Unwrap() error { return er.err }

func (er *AppError) Marshal() []byte {
	resp := errResponce{Code: er.Code, Err: er.err.Error()}
	bytes, err := json.Marshal(&resp)
	if err != nil {
		return nil
	}
	return bytes
}
func BadRequestErr(loc string, err error) *AppError {
	appErr := &AppError{Code: ErrBadRequest, Location: loc, err: err}
	return appErr
}
func InternalErr(loc string, err error) *AppError {
	appErr := &AppError{Code: ErrInternal, Location: loc, err: err}
	return appErr
}
func NotFoundErr(loc string, err error) *AppError {
	appErr := &AppError{Code: ErrNotFound, Location: loc, err: err}
	return appErr
}

type errResponce struct {
	Code errorCode `json:"code"`
	Err  string    `json:"error"`
}
