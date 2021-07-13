package services

import (
	"github.com/deepak-v4/bookstore_users-api/domain/users"
	"github.com/deepak-v4/bookstore_users-api/utils/error"
)

func CreateUser(user users.User) (*users.User, *error.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userid int64) (*users.User, *error.RestErr) {
	if userid <= 0 {
		return nil, error.NewBadRequest("invalid userid")
	}

	result := &users.User{Id: userid}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil

}
