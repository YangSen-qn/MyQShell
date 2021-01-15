package common

import (
	"encoding/json"
)

const (
	ErrorCodeCMDRegisterInvalidParams = -10000
)

type QShellError struct {
	code int `json:"code"`
	description string `json:"description"`
}

func NewQShellError(code int, description string) *QShellError	  {
	return &QShellError{
		code:        code,
		description: description,
	}
}

func (err *QShellError)MarshalJSON() ([]byte, error) {
	return json.Marshal(err)
}

func (e *QShellError) Error() string {
	return e.description
}

type IError interface {
	error
	json.Marshaler

	ErrorCode() int
}
