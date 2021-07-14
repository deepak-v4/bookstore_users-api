package services

import (
	"github.com/deepak-v4/bookstore_users-api/domain/users"
	"github.com/deepak-v4/bookstore_users-api/utils/date"
	"github.com/deepak-v4/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	user.DateCreated = date.GetTimeString()

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userid int64) (*users.User, *errors.RestErr) {
	if userid <= 0 {
		return nil, errors.NewBadRequest("invalid userid")
	}

	result := &users.User{Id: userid}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil

}
