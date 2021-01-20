package utils

import (
	"encoding/base64"
	"strings"
)

// 保存在account.json文件中的数据格式
func Encrypt(accessKey, encryptedKey, name string) string {
	return strings.Join([]string{name, accessKey, encryptedKey}, ":")
}

// 对SecretKey加密, 返回加密后的字符串
func EncryptSecretKey(accessKey, secretKey string) (string, error) {
	aesKey := Md5Hex(accessKey)
	encryptedSecretKeyBytes, encryptedErr := AesEncrypt([]byte(secretKey), []byte(aesKey[7:23]))
	if encryptedErr != nil {
		return "", encryptedErr
	}
	encryptedSecretKey := base64.URLEncoding.EncodeToString(encryptedSecretKeyBytes)
	return encryptedSecretKey, nil
}

// 对加密的SecretKey进行解密， 返回SecretKey
func DecryptSecretKey(accessKey, encryptedKey string) (string, error) {
	aesKey := Md5Hex(accessKey)
	encryptedSecretKeyBytes, decodeErr := base64.URLEncoding.DecodeString(encryptedKey)
	if decodeErr != nil {
		return "", decodeErr
	}
	secretKeyBytes, decryptErr := AesDecrypt([]byte(encryptedSecretKeyBytes), []byte(aesKey[7:23]))
	if decryptErr != nil {
		return "", decryptErr
	}
	secretKey := string(secretKeyBytes)
	return secretKey, nil
}