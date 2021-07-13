package users

import (
	"fmt"

	"github.com/deepak-v4/bookstore_users-api/utils/error"
)

var (
	UsersDB = make(map[int64]*User)
)

func (user *User) Get() *error.RestErr {

	result := UsersDB[user.Id]
	if result == nil {
		return error.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *error.RestErr {

	current := UsersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return error.NewBadRequest(fmt.Sprintf("User %s email already registered", user.Email))
		}
		return error.NewBadRequest(fmt.Sprintf("User %d already exist", user.Id))
	}

	UsersDB[user.Id] = user
	return nil
}
