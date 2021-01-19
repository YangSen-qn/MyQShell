package config

import "github.com/spf13/viper"

var (
	keyCredentialDBPath = []string{"path.credential_db", "path.credential_db_path"}
	keyCurrentCredentialName = "currentCredentialName"

	credentialDBPath = ""
	currentCredentialName = ""
)

func GetCredentialDBPath() string {
	if credentialDBPath != "" {
		return credentialDBPath
	}

	path := viper.GetString(keyCredentialDBPath[0])
	if path != "" {
		return path
	}
	return viper.GetString(keyCredentialDBPath[1])
}

func SetCredentialDBPath(path string) {
	credentialDBPath = path

	for _, key := range keyCredentialDBPath {
		viper.Set(key, path)
	}
}

func GetCurrentCredentialName() string {
	return currentCredentialName
}

func SetCurrentCredentialName(name string) {
	currentCredentialName = name
}

