package load

import (
	"qshell/iqshell/config"
	"qshell/iqshell/user"
	"qshell/qn_shell_error"
)

func LoadInterQShell(loadConfig *config.LoadConfig) qn_shell_error.IQShellError {
	config.SetLoadConfig(loadConfig)

	dbPath, err := config.GetCredentialDBPath()
	if err != nil {
		return err
	}
	user.SetDBPath(dbPath)

	configPath, err := config.GetConfigPath()
	err = user.SetCachePath(configPath)
	if err != nil {
		return err
	}

	return nil
}
