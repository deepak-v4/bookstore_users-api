package users

import (
	"fmt"

	"github.com/deepak-v4/bookstore_users-api/datasources/mysql/user_db"
	"github.com/deepak-v4/bookstore_users-api/utils/errors"
	"github.com/deepak-v4/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	querySelectUser = "SELECT id,first_name,last_name,email,date_created from users where id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
)

var (
	UsersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	stmt, err := user_db.Client.Prepare(querySelectUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	searchResult := stmt.QueryRow(user.Id)

	if sel_err := searchResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); sel_err != nil {
		return mysql_utils.ParseError(sel_err)
		//return errors.NewNotFoundError(fmt.Sprintf("When trying to fetch user details %s", sel_err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveerr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveerr != nil {
		return mysql_utils.ParseError(saveerr)
		//return errors.NewInternalServerError("internal server error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError("internal server error")
	}

	user.Id = userId

	return nil

}

func (user *User) Update() *errors.RestErr {

	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	fmt.Println("Hello alive", user)
	_, saveerr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if saveerr != nil {
		fmt.Println("hello alive 888", saveerr)
		return mysql_utils.ParseError(saveerr)
		//return errors.NewInternalServerError("internal server error")
	}
	fmt.Println("Hello alive 7777")
	return nil
}
