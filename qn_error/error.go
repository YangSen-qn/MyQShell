package qn_error

import (
	"encoding/json"
	"fmt"
)

type qnError struct {
	err         error
	Description string `json:"error"`
}

func NewError(format string, a ...interface{}) error {
	return &qnError{
		err:         nil,
		Description: fmt.Sprintf(format, a...),
	}
}

func NewErrorWithError(err error) error {
	return NewErrorWithErrorFormat(err, "")
}

func NewErrorWithErrorFormat(err error, format string, a ...interface{}) error {
	return &qnError{
		err:         err,
		Description: fmt.Sprintf(format, a...),
	}
}

func (err *qnError) String() string {
	desc := err.Description
	if err.err != nil {
		desc = fmt.Sprintf("***%s[%s]", desc, err.err)
	}
	return desc
}

func (err *qnError) marshalObject() error {
	return NewError(err.String())
}

func (err *qnError) Error() string {
	return err.String()
}

func ToJson(err error) string {
	qnErr, ok := err.(*qnError)
	if !ok {
		qnErr, _ = NewErrorWithError(err).(*qnError)
	}

	data, err := json.Marshal(qnErr.marshalObject())
	if err == nil {
		return string(data)
	} else {
		return ""
	}
}