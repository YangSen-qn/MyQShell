package user

import (
	"qshell/qn_shell_error"
)

var (
	currentCredential *Credential = nil
)

type Credential struct {
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (credential *Credential) String() string {
	return "Name:" + credential.Name + " AccessKey:" + credential.AccessKey + " SecretKey:" + credential.SecretKey
}

func (credential *Credential) isValid() bool {
	return credential.Name != "" && credential.AccessKey != "" && credential.SecretKey != ""
}

func (credential *Credential) checkValid() qn_shell_error.IQShellError {
	if credential.isValid() {
		return nil
	} else {
		return qn_shell_error.NewInvalidUserParamError("credential is invalid for name:" + credential.Name)
	}
}

// ------------- api -------------
func CredentialList() [] *Credential {
	return credentialListFromDB()
}

func CurrentCredential() *Credential {
	if currentCredential != nil {
		return currentCredential
	}

	currentCredential = cacheGetCurrentCredential()
	if currentCredential != nil && currentCredential.isValid() {
		return currentCredential
	}

	allCredential := CredentialList()
	for _, credential := range allCredential {
		if credential.isValid() {
			currentCredential = credential
			break
		}
	}

	if currentCredential.isValid() {
		cacheSetCurrentCredential(currentCredential)
		return currentCredential
	} else {
		return nil
	}
}

func SetCurrentCredential(name string) qn_shell_error.IQShellError {
	credential := GetCredential(name)
	if credential == nil {
		return qn_shell_error.NewInvalidUserParamError("not exist credential for name:" + name)
	}

	if !credential.isValid() {
		return qn_shell_error.NewInvalidUserParamError("credential is invalid for name:" + name)
	}

	currentCredential = credential
	cacheSetCurrentCredential(currentCredential)
	return nil
}

func GetCredential(name string) *Credential {
	return getCredentialFromDB(name)
}

func RemoveCredential(name string) qn_shell_error.IQShellError {
	err := removeCredentialFromDB(name)
	if err != nil {
		return err
	}

	if currentCredential != nil && name == currentCredential.Name {
		currentCredential = nil
		cacheRemoveCurrentCredential()
	}

	return err
}

func AddCredential(credential *Credential) qn_shell_error.IQShellError {
	if !credential.isValid() {
		return qn_shell_error.NewInvalidUserParamError("credential info is invalid")
	}

	cacheSetCurrentCredential(credential)

	return addCredentialToDB(credential, false)
}

