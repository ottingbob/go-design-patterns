package proxy

import (
	"fmt"
)

// Wraps an object to hide some of its characteristics
// Reasons for restricting access could be:
// 1) Remote object (remote proxy)
// 2) Large chunk of memory (virtual proxy)
// 3) Restricted access object (protection proxy)

// Provides a new abstraction layer that is easy to work
// with and may be changed easily

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}

	return User{}, fmt.Errorf("User %s could not be found\n", id)
}

type UserListProxy struct {
	SomeDatabase  		  *UserList
	StackCache 	  		  UserList
	StackCapacity 		  int
	DidLastSearchUseCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUseCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.DidLastSearchUseCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}