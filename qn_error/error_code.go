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

func NewDBError(description string) *Error {
	return NewHeavyError(QShellErrorCodeDBPathError, description)
}

func NewFilePathError(description string) *Error {
	return NewHeavyError(QShellErrorCodeFilePathError, description)
}

func NewCryptError(description string) *Error {
	return NewHeavyError(QShellErrorCodeCryptError, description)
}

func NewIOError(description string) *Error {
	return NewHeavyError(QShellErrorCodeIOError, description)
}

func NewWarningError(description string) *Error {
	return NewHeavyError(QShellErrorCodeWarning, description)
}