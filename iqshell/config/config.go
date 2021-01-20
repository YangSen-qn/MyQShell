package config

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"qshell/iqshell/cache"
	"qshell/qn_shell_error"
)

const (
	configFileName = "config.json"

	// 根路径名
	rootDirName = ".qshell"
	// credential db 文件名
	credentialDBName = "account.db"
)

var (
	// 配置信息的key
	// Credential path
	keyCredentialDBPath = []string{"path.accdb", "path.acc_db_path"}

	// 根路径
	config *LoadConfig = nil

	// 配置缓存
	configCache = cache.NewCache()
)

type LoadConfig struct {
	// 根路径
	RootDir string
}


func SetLoadConfig(loadConfig *LoadConfig)  {
	config = loadConfig
}

func GetCredentialDBPath() (path string, err qn_shell_error.IQShellError) {
	path, err = RootPath()
	if err != nil {
		return
	}

	path = filepath.Join(path, credentialDBName)
	return
}

// config path
func GetConfigPath() (path string, err qn_shell_error.IQShellError)  {
	path, err = RootPath()
	if err != nil {
		return
	}

	path = path + "/" + configFileName
	return
}

// 获取ROOTPath
func RootPath() (path string, err qn_shell_error.IQShellError) {
	if config.RootDir == "" {
		config.RootDir = defaultRootPath()
	}
	if config.RootDir == "" {
		err = qn_shell_error.NewInvalidFilePathError("root dir not exist")
	}
	path = config.RootDir
	return
}

func defaultRootPath() string {
	dir, err := homedir.Dir()
	if err == nil {
		return filepath.Join(dir, rootDirName)
	} else {
		return ""
	}
}

func LocalPath() string {
	path, _ := os.Getwd()
	return path
}
