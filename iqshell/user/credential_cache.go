package user

import (
	"qshell/iqshell/cache"
)

var (
	keyCurrentCredentialName = []string{"current_credential_name"}
	keyCurrentCredentialAK   = []string{"current_credential_ak"}
	keyCurrentCredentialSK   = []string{"current_credential_sk"}

	// 配置缓存
	currentCredentialCache = cache.NewCache()
)

func setCachePath(path string) {
	if path != "" {
		currentCredentialCache.SetCacheFile(path)
	}
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

