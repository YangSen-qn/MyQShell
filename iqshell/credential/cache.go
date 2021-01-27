package credential

import (
	"qshell/iqshell/cache"
	"qshell/util"
)

var (
	keyCurrentCredentialName = []string{"current_credential_name"}
	keyCurrentCredentialAK   = []string{"current_credential_ak"}
	keyCurrentCredentialSK   = []string{"current_credential_sk"}

	// 配置缓存
	currentCredentialCache = cache.NewCache()
)

func SetCachePath(path string) error {
	if !util.Exist(path) || util.IsFileEmpty(path) {
		util.CreateFile(path, "{}")
	}
	return currentCredentialCache.SetCacheFile(path)
}

func cacheRemoveCurrentCredential() {
	currentCredentialCache.CacheSetString("", keyCurrentCredentialName)
	currentCredentialCache.CacheSetString("", keyCurrentCredentialAK)
	currentCredentialCache.CacheSetString("", keyCurrentCredentialSK)
}

func cacheSetCurrentCredential(credential *Credential) {
	currentCredentialCache.CacheSetString(credential.Name, keyCurrentCredentialName)
	currentCredentialCache.CacheSetString(credential.AccessKey, keyCurrentCredentialAK)
	currentCredentialCache.CacheSetString(credential.SecretKey, keyCurrentCredentialSK)
}

func cacheGetCurrentCredential() *Credential {
	return &Credential{
		Name:      currentCredentialCache.CacheGetString(keyCurrentCredentialName),
		AccessKey: currentCredentialCache.CacheGetString(keyCurrentCredentialAK),
		SecretKey: currentCredentialCache.CacheGetString(keyCurrentCredentialSK),
	}
}
