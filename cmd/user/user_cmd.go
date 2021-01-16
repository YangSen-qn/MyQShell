package user

import (
	"qshell/execute"
	"qshell/iqshell/user"
	"qshell/output"
)

// ------------- List ---------------
type userListCMD struct {
	execute.Command
}

func (cmd *userListCMD) Check() error {
	return nil
}

func (cmd *userListCMD) Execute() error {
	userList := user.UserList()

	u := &user.User{
		IsCurrent: false,
		Name:      "kodo",
		AccessKey: "accessKey",
		SecretKey: "secretKey",
	}
	output.OutputResult(cmd, u)

	for _, u := range userList {
		output.OutputResult(cmd, u)
	}
	return nil
}

// ------------- Get ---------------
type userGetCMD struct {
	execute.Command

	name string
}

func (cmd *userGetCMD) Check() error {
	return nil
}

func (cmd *userGetCMD) Execute() error {
	return nil
}

// ------------- Add ---------------
type userAddCMD struct {
	execute.Command

	name      string
	accessKey string
	secretKey string
}

func (cmd *userAddCMD) Check() error {
	return nil
}

func (cmd *userAddCMD) Execute() error {
	return nil
}

// ------------- Remove ---------------
type userRemoveCMD struct {
	execute.Command

	name string
}

func (cmd *userRemoveCMD) Check() error {
	return nil
}

func (cmd *userRemoveCMD) Execute() error {
	return nil
}

// ------------- currentGet ---------------
type currentUserGetCMD struct {
	execute.Command
}

func (cmd *currentUserGetCMD) Check() error {
	return nil
}

func (cmd *currentUserGetCMD) Execute() error {
	return nil
}

// ------------- CurrentSet ---------------
type currentUserSetCMD struct {
	execute.Command

	name string
}

func (cmd *currentUserSetCMD) Check() error {
	return nil
}

func (cmd *currentUserSetCMD) Execute() error {
	return nil
}
