package qn_error

import (
	"encoding/json"
)

type Level int8

const (
	LevelLight Level = iota
	LevelMiddle
	LevelHeavy
)

type IError interface {
	error
	json.Marshaler

	ErrorLevel() Level
	ErrorCode() string
}

type Error struct {
	level       Level
	code        string `json:"code"`
	description string `json:"description"`
}

func NewError(level Level, code string, description string) *Error {
	return &Error{
		level:       level,
		code:        code,
		description: description,
	}
}

func NewLightError(code string, description string) *Error {
	return NewError(LevelLight, code,description)
}

func NewMiddleError(code string, description string) *Error {
	return NewError(LevelMiddle, code,description)
}

func NewHeavyError(code string, description string) *Error {
	return NewError(LevelHeavy, code,description)
}

func (err *Error) String() string {
	return err.code + ": " + err.description
}

func (err *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(err)
}

func (err *Error) Error() string {
	return err.code + ": " + err.description
}

func (err *Error) ErrorCode() string {
	return err.code
}

func (err *Error) ErrorLevel() Level {
	return err.level
}