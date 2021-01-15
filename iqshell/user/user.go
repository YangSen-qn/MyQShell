package user

type User struct {
	IsCurrent bool   `json:"is_current"`
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (user *User) String() string {
	return "Name:" + user.Name + " AccessKey:" + user.AccessKey + " SecretKey:" + user.SecretKey
}

// ------------- api -------------
func init() {
	list := userListFromCache()
	for _, user := range list {
		_ = addUserToCache(user)
	}
}

func UserList() [] *User {
	list := userListFromCache()
	if list != nil {
		return list
	}

	list = userListFromDB()
	return list
}

func CurrentUser(isGlobal bool) *User {
	if isGlobal {
		return currentUserFromCache()
	} else {
		return currentUserFromDB()
	}
}

func SetCurrentUser(userName string, isGlobal bool) error {
	if isGlobal {
		return setCurrentUserToCache(userName)
	} else {
		return setCurrentUserToDB(userName)
	}
}

func GetUser(userName string) *User {
	user := getUserFromCache(userName)
	if user != nil {
		return user
	}

	user = getUserFromDB(userName)
	return user
}

func RemoveUser(userName string) error {
	err := removeUserToCache(userName)
	err = removeUserToDB(userName)
	return err
}

func AddUser(user *User) error {
	err := addUserToDB(user)
	err = addUserToCache(user)
	return err
}
