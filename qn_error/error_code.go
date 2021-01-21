package qn_error

const (
	QShellErrorCodeWarning = "Warning"
	QShellErrorCodeIOError = "IO error"
	QShellErrorCodeCryptError = "crypt error"
	QShellErrorCodeDBPathError = "db error"
	QShellErrorCodeFilePathError = "filePath error"
	QShellErrorCodeInvalidUserParam = "invalid user input param"
)


func NewInvalidUserParamError(description string) *Error {
	return NewHeavyError(QShellErrorCodeInvalidUserParam, description)
}

func NewInvalidDBError(description string) *Error {
	return NewHeavyError(QShellErrorCodeDBPathError, description)
}

func NewInvalidFilePathError(description string) *Error {
	return NewHeavyError(QShellErrorCodeFilePathError, description)
}

func NewInvalidCryptError(description string) *Error {
	return NewHeavyError(QShellErrorCodeCryptError, description)
}

func NewInvalidIOError(description string) *Error {
	return NewHeavyError(QShellErrorCodeIOError, description)
}

func NewInvalidWarningError(description string) *Error {
	return NewHeavyError(QShellErrorCodeWarning, description)
}