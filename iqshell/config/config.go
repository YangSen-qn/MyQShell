package config

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"qshell/iqshell/cache"
	"qshell/qn_shell_error"
)

const (
	defaultConfigFileName = ".qshell"

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
	rootPath = ""

	// 配置缓存
	configCache = cache.NewCache()
)

func init() {
	SetConfigPath("")
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
func GetConfigPath() string {
	return configCache.GetCachePath()
}

func SetConfigPath(path string) qn_shell_error.IQShellError {
	if path != "" {
		return configCache.SetCacheFile(path)
	}

	homeDir, err := RootPath()
	if err != nil {
		return err
	} else {
		return configCache.SetCachePath(homeDir, defaultConfigFileName)
	}
}

// 获取ROOTPath
func RootPath() (path string, err qn_shell_error.IQShellError) {
	if rootPath == "" {
		SetRootPath(defaultRootPath())
	}
	path = rootPath
	if path == "" {
		err = qn_shell_error.NewInvalidFilePathError("root dir not exist")
	}
	return
}

// 设置RootPath
func SetRootPath(path string) {
	rootPath = path
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
