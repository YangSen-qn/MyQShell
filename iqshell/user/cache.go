package user

import (
	"qshell/common"
)

var (
	currentUser *User = nil
	userCache         = make(map[string]*User)
)

func userListFromCache() [] *User {
	userList := make([]*User, 0)
	for _, user := range userCache {
		userList = append(userList, user)
	}
	return userList
}

func currentUserFromCache() *User {
	return currentUser
}

func setCurrentUserToCache(userName string) error {
	user := userCache[userName]
	if user == nil {
		return &common.QShellError{}
	}
	return nil
}

func getUserFromCache(userName string) *User {

	return nil
}

func removeUserToCache(userName string) error {

	return nil
}

func addUserToCache(user *User) error {

	return nil
}
