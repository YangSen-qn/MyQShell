package credential

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"qshell/qn_error"

	"github.com/syndtr/goleveldb/leveldb"
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

func removeCredentialFromDB(name string) qn_error.IError {
	if !isDBPathValid(dbPath) {
		return qn_error.NewFilePathError("db path is invalid")
	}

	db, dbErr := leveldb.OpenFile(dbPath, nil)
	if dbErr != nil {
		return qn_error.NewDBError(dbErr.Error())
	}
	defer db.Close()

	dbWOpt := &opt.WriteOptions{
		Sync: true,
	}
	dbErr = db.Delete([]byte(name), dbWOpt)
	if dbErr != nil {
		return qn_error.NewDBError(dbErr.Error())
	}

	return nil
}

func addCredentialToDB(credential *Credential, isCover bool) qn_error.IError {
	if !isDBPathValid(dbPath) {
		return qn_error.NewFilePathError("db path is invalid")
	}

	db, dbErr := leveldb.OpenFile(dbPath, nil)
	if dbErr != nil {
		return qn_error.NewDBError(dbErr.Error())
	}
	defer db.Close()

	if !isCover {
		exists, dbErr := db.Has([]byte(credential.Name), nil)
		if dbErr != nil {
			return qn_error.NewDBError(dbErr.Error())
		}

		if exists {
			return qn_error.NewInvalidUserParamError("credential name:" + credential.Name + " already exist in local db")
		}
	}

	dbWOpt := &opt.WriteOptions{
		Sync: true,
	}

	credential, err := encrypt(credential)
	if err != nil {
		return err
	}

	dbKey := []byte(credentialDBId(credential))
	dbValue, err := credentialToDBValue(credential)
	if err != nil {
		return err
	}

	dbErr = db.Put(dbKey, dbValue, dbWOpt)
	if dbErr != nil {
		return qn_error.NewDBError(dbErr.Error())
	}

	return nil
}

// 对SecretKey进行加密， 保存AccessKey, 加密后的SecretKey在本地数据库中
func encrypt(credential *Credential) (newCredential *Credential, err qn_error.IError) {
	encryptedKey, eErr := encryptSecretKey(credential.AccessKey, credential.SecretKey)
	if eErr != nil {
		err = qn_error.NewCryptError("secret key encrypt error")
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
func decrypt(credential *Credential) (newCredential *Credential, err qn_error.IError) {
	secretKey, dErr := decryptSecretKey(credential.AccessKey, credential.SecretKey)
	if dErr != nil {
		err = qn_error.NewCryptError("secret key decrypt error")
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

func credentialToDBValue(credential *Credential) (value []byte, err qn_error.IError) {
	if err != nil {
		return
	}

	value, jErr := json.Marshal(credential)
	if jErr != nil {
		err = qn_error.NewIOError(jErr.Error())
	}
	return
}

func credentialFromDBValue(value []byte) (credential *Credential, err qn_error.IError) {
	credential = &Credential{}
	jErr := json.Unmarshal(value, credential)
	if jErr != nil {
		err = qn_error.NewIOError(jErr.Error())
	}
	return
}