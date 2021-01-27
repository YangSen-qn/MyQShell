package qn_error

import (
	"encoding/json"
	"fmt"
)

type qnError struct {
	err         error  `json:"error_info"`
	description string `json:"error"`
}

func NewError(format string, a ...interface{}) error {
	return &qnError{
		err:         nil,
		description: fmt.Sprintf(format, a),
	}
}

func NewErrorWithError(err error) error {
	return NewErrorWithErrorFormat(err, "")
}

func NewErrorWithErrorFormat(err error, format string, a ...interface{}) error {
	return &qnError{
		err:         err,
		description: fmt.Sprintf(format, a),
	}
}

func (err *qnError) String() string {
	return "error: " + err.description
}

func (err *qnError) MarshalJSON() ([]byte, error) {
	return json.Marshal(err)
}

func (err *qnError) Error() string {
	return "error: " + err.description
}

func ToJson(err error) string {
	qnErr, ok := err.(*qnError)
	if ok == false {
		qnErr, _ = NewErrorWithError(err).(*qnError)
	}

	data, err := json.Marshal(qnErr)
	if err == nil {
		return string(data)
	} else {
		return ""
	}
}