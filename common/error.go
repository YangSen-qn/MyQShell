package common

import (
	"encoding/json"
)

const (
	ErrorCodeCMDRegisterInvalidParams = -10000
)

type IQShellError interface {
	error
	json.Marshaler

	ErrorCode() int
}

type QShellError struct {
	code        int    `json:"code"`
	description string `json:"description"`
}

func NewQShellError(code int, description string) *QShellError {
	return &QShellError{
		code:        code,
		description: description,
	}
}

func (err *QShellError) MarshalJSON() ([]byte, error) {
	return json.Marshal(err)
}

func (err *QShellError) Error() string {
	return err.description
}

func (err *QShellError) ErrorCode() int {
	return err.code
}
