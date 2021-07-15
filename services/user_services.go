package services

import (
	"github.com/deepak-v4/bookstore_users-api/domain/users"
	cryptoutils "github.com/deepak-v4/bookstore_users-api/utils/crypto_utils"
	"github.com/deepak-v4/bookstore_users-api/utils/date"
	"github.com/deepak-v4/bookstore_users-api/utils/errors"
)

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := GetUser(int64(user.Id))
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {

		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	user.DateCreated = date.GetTimeString()
	user.Password = cryptoutils.GetMd5(user.Password)
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

func DeleteUser(userid int64) *errors.RestErr {

	if userid <= 0 {
		return errors.NewBadRequest("invalid userid")
	}

	result := &users.User{Id: userid}
	if err := result.Delete(); err != nil {
		return err
	}

	return nil
}

func SearchByStatus(usr_status string) ([]users.User, *errors.RestErr) {

	dao := &users.User{}
	users, err := dao.FindByStatus(usr_status)
	if err != nil {
		return nil, err
	}
	return users, nil
}
