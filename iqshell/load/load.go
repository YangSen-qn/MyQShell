package load

import (
	"qshell/iqshell/config"
	"qshell/iqshell/credential"
)

func LoadInterQShell(loadConfig *config.LoadConfig) error {

	// 加在配置
	config.SetLoadConfig(loadConfig)

	dbPath, err := config.GetCredentialDBPath()
	if err != nil {
		return err
	}
	credential.SetDBPath(dbPath)

	configPath, err := config.GetConfigPath()
	err = credential.SetCachePath(configPath)
	if err != nil {
		return err
	}

	return nil
}
