package qn_shell_error

const (
	QShellErrorCodeWarning = "Warning"
	QShellErrorCodeIOError = "IO error"
	QShellErrorCodeCryptError = "crypt error"
	QShellErrorCodeDBPathError = "db error"
	QShellErrorCodeFilePathError = "filePath error"
	QShellErrorCodeInvalidUserParam = "invalid user input param"
)


func NewInvalidUserParamError(description string) *QShellError {
	return NewHeavyError(QShellErrorCodeInvalidUserParam, description)
}

func NewInvalidDBError(description string) *QShellError {
	return NewHeavyError(QShellErrorCodeDBPathError, description)
}

func NewInvalidFilePathError(description string) *QShellError {
	return NewHeavyError(QShellErrorCodeFilePathError, description)
}

func NewInvalidCryptError(description string) *QShellError {
	return NewHeavyError(QShellErrorCodeCryptError, description)
}

func NewInvalidIOError(description string) *QShellError {
	return NewHeavyError(QShellErrorCodeIOError, description)
}

func NewInvalidWarningError(description string) *QShellError {
	return NewHeavyError(QShellErrorCodeWarning, description)
}