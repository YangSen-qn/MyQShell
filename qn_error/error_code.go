package qn_error

const (
	QShellErrorCodeExecuteError = "execute error: "
	QShellErrorCodeFilePathError = "filePath error: "
	QShellErrorCodeInvalidUserParam = "invalid user input param: "
)

func NewInvalidUserParamError(format string, a ...interface{}) error {
	format = QShellErrorCodeInvalidUserParam + format
	return NewError(format, a...)
}

func NewFilePathError(format string, a ...interface{}) error {
	format = QShellErrorCodeFilePathError + format
	return NewError(format, a...)
}

func NewExecuteError(format string, a ...interface{}) error {
	format = QShellErrorCodeExecuteError + format
	return NewError(format, a...)
}
