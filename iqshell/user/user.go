package user

import "qshell/common"

const (
	CurrentCredentialNameKey = "currentCredential"
)

var (
	currentCredential *Credential = nil
)

type Credential struct {
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (user *Credential) String() string {
	return "Name:" + user.Name + " AccessKey:" + user.AccessKey + " SecretKey:" + user.SecretKey
}

// ------------- api -------------
func init() {
	setupCurrentCredential()
}

func getCurrentCredentialNameFromDisk() string {
	return ""
}

func saveCurrentCredentialNameToDisk() {

}

func setupCurrentCredential() {
	currentCredential = CurrentCredential()
	if currentCredential != nil {
		return
	}

	allCredential := CredentialList()
	if allCredential != nil && len(allCredential) > 0 {
		currentCredential = allCredential[0]
	} else {
		currentCredential = nil
	}

	if currentCredential != nil {
		saveCurrentCredentialNameToDisk()
	}
}

func CredentialList() [] *Credential {
	return credentialListFromDB()
}

func CurrentCredential() *Credential {
	if currentCredential != nil {
		return currentCredential
	}
	currentCredential = GetCredential(getCurrentCredentialNameFromDisk())
	return currentCredential
}

func SetCurrentCredential(name string) common.IQShellError {
	c := GetCredential(name)
	if c == nil {
		return common.NewQShellError(-1, "not exist credential for name:"+name)
	}

	currentCredential = c
	saveCurrentCredentialNameToDisk()
	return nil
}

func GetCredential(name string) *Credential {
	return getCredentialFromDB(name)
}

func RemoveCredential(name string) common.IQShellError {
	err := removeCredentialFromDB(name)
	if err != nil {
		return err
	}

	if name == currentCredential.Name {
		currentCredential = nil
		setupCurrentCredential()
	}

	return err
}

func AddCredential(credential *Credential) common.IQShellError {
	err := addCredentialToDB(credential)
	if err  == nil{
		setupCurrentCredential()
	}
	return err
}
