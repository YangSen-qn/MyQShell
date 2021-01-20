package qn_shell_error

import (
	"encoding/json"
)

type QShellErrorLevel int8

const (
	QShellErrorLevelLight QShellErrorLevel = iota
	QShellErrorLevelMiddle
	QShellErrorLevelHeavy
)

type IQShellError interface {
	error
	json.Marshaler

	ErrorLevel() QShellErrorLevel
	ErrorCode() string
}

type QShellError struct {
	level       QShellErrorLevel
	code        string `json:"code"`
	description string `json:"description"`
}

func NewError(level QShellErrorLevel, code string, description string) *QShellError {
	return &QShellError{
		level:       level,
		code:        code,
		description: description,
	}
}

func NewLightError(code string, description string) *QShellError {
	return NewError(QShellErrorLevelLight, code,description)
}

func NewMiddleError(code string, description string) *QShellError {
	return NewError(QShellErrorLevelMiddle, code,description)
}

func NewHeavyError(code string, description string) *QShellError {
	return NewError(QShellErrorLevelHeavy, code,description)
}

func (err *QShellError) String() string {
	return err.code + ": " + err.description
}

func (err *QShellError) MarshalJSON() ([]byte, error) {
	return json.Marshal(err)
}

func (err *QShellError) Error() string {
	return err.code + ": " + err.description
}

func (err *QShellError) ErrorCode() string {
	return err.code
}

func (err *QShellError) ErrorLevel() QShellErrorLevel {
	return err.level
}