package users

import (
	"strings"

	"github.com/deepak-v4/bookstore_users-api/utils/error"
)

type User struct {
	Id          int64  `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *error.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return error.NewBadRequest("invalid email address")
	}
	return nil
}
