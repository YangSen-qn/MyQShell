package cache

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"qshell/qn_shell_error"
)

type Cache struct {
	cache          *viper.Viper
	configJsonFile string
}

func NewCache() *Cache {
	return &Cache{
		cache:          viper.New(),
		configJsonFile: "",
	}
}

func (cache *Cache) GetCachePath() string {
	return cache.configJsonFile
}

func (cache *Cache) SetCacheFile(file string) (err qn_shell_error.IQShellError) {
	if file == "" {
		err = qn_shell_error.NewInvalidFilePathError("cache file is empty")
		return
	} else {
		cache.configJsonFile = file
		cache.cache.SetConfigFile(cache.configJsonFile)
	}

	return cache.ReadInConfig()
}

func (cache *Cache) SetCachePath(path string, name string) qn_shell_error.IQShellError {
	if path == "" {
		return qn_shell_error.NewInvalidFilePathError("cache path is empty")
	} else {
		cache.cache.AddConfigPath(path)
		cache.cache.SetConfigName(name)
		cache.configJsonFile = filepath.Join(path, name)
	}

	return cache.ReadInConfig()
}

func (cache *Cache) ReadInConfig() qn_shell_error.IQShellError {
	err := cache.cache.ReadInConfig()
	if err == nil {
		return nil
	} else {
		return qn_shell_error.NewInvalidFilePathError(err.Error())
	}
}

// api
func (cache *Cache) CacheSetString(value string, keyList []string) {
	for _, key := range keyList {
		fmt.Println("write config key:", key, " value:", value)
		cache.cache.Set(key, value)
		cache.cache.WriteConfig()
	}
}

func (cache *Cache) CacheGetString(keyList []string) string {
	value := ""
	for _, key := range keyList {
		if key != "" {
			value = cache.cache.GetString(key)
			break
		}
	}
	return value
}