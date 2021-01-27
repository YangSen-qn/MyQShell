package credential

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"qshell/qn_error"
)

var (
	dbPath string
)

func SetDBPath(path string) {
	dbPath = path
}

func isDBPathValid(path string) bool {
	return path != ""
}

// 全局有效
func credentialListFromDB() [] *Credential {
	if !isDBPathValid(dbPath) {
		return nil
	}

	db, dbErr := leveldb.OpenFile(dbPath, nil)
	if dbErr != nil {
		return nil
	}
	defer db.Close()

	credentialList := make([] *Credential, 0)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		credential, dbErr := credentialFromDBValue(iter.Value())
		if dbErr != nil {
			continue
		}

		credential, err := decrypt(credential)
		if err != nil {
			continue
		}

		if credential.isValid() {
			credentialList = append(credentialList, credential)
		}
	}
	iter.Release()

	return credentialList
}

func getCredentialFromDB(name string) *Credential {
	if !isDBPathValid(dbPath) {
		return nil
	}

	db, dbErr := leveldb.OpenFile(dbPath, nil)
	if dbErr != nil {
		return nil
	}
	defer db.Close()

	value, dbErr := db.Get([]byte(name), nil)
	if dbErr != nil {
		return nil
	}

	credential, dbErr := credentialFromDBValue(value)
	if dbErr != nil {
		return nil
	}

	credential, err := decrypt(credential)
	if err != nil {
		return nil
	}

	return credential
}

func removeCredentialFromDB(name string) error {
	if !isDBPathValid(dbPath) {
		return qn_error.NewFilePathError("db path is invalid")
	}

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	dbWOpt := &opt.WriteOptions{
		Sync: true,
	}
	err = db.Delete([]byte(name), dbWOpt)
	if err != nil {
		return err
	}

	return nil
}

func addCredentialToDB(credential *Credential, isCover bool) error {
	if !isDBPathValid(dbPath) {
		return qn_error.NewFilePathError("db path is invalid")
	}

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	if !isCover {
		exists, err := db.Has([]byte(credential.Name), nil)
		if err != nil {
			return err
		}

		if exists {
			return qn_error.NewInvalidUserParamError("credential name: %s already exist in local db", credential.Name)
		}
	}

	dbWOpt := &opt.WriteOptions{
		Sync: true,
	}

	credential, err = encrypt(credential)
	if err != nil {
		return err
	}

	dbKey := []byte(credentialDBId(credential))
	dbValue, err := credentialToDBValue(credential)
	if err != nil {
		return err
	}

	err = db.Put(dbKey, dbValue, dbWOpt)
	if err != nil {
		return err
	}

	return nil
}

// 对SecretKey进行加密， 保存AccessKey, 加密后的SecretKey在本地数据库中
func encrypt(credential *Credential) (newCredential *Credential, err error) {
	encryptedKey, err := encryptSecretKey(credential.AccessKey, credential.SecretKey)
	if err != nil {
		return
	}

	newCredential = &Credential{
		Name:      credential.Name,
		AccessKey: credential.AccessKey,
		SecretKey: encryptedKey,
	}
	return
}

// 对本地数据库中 credential 进行解密
func decrypt(credential *Credential) (newCredential *Credential, err error) {
	secretKey, err := decryptSecretKey(credential.AccessKey, credential.SecretKey)
	if err != nil {
		return
	}

	newCredential = &Credential{
		Name:      credential.Name,
		AccessKey: credential.AccessKey,
		SecretKey: secretKey,
	}
	return
}

func credentialDBId(credential *Credential) string {
	return credential.Name
}

func credentialToDBValue(credential *Credential) (value []byte, err error) {
	if err != nil {
		return
	}

	value, err = json.Marshal(credential)
	return
}

func credentialFromDBValue(value []byte) (credential *Credential, err error) {
	credential = &Credential{}
	err = json.Unmarshal(value, credential)
	return
}