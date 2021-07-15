package users

import (
	"github.com/deepak-v4/bookstore_users-api/datasources/mysql/user_db"
	"github.com/deepak-v4/bookstore_users-api/utils/errors"
	"github.com/deepak-v4/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created,status,Password) VALUES(?,?,?,?,?,?);"
	querySelectUser = "SELECT id,first_name,last_name,email,date_created,status from users where id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users where id=?;"
	querySearchUser = "SELECT id,first_name,last_name,email,date_created,status from users where status=?;"
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

	if sel_err := searchResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); sel_err != nil {
		return mysql_utils.ParseError(sel_err)

	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveerr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveerr != nil {
		return mysql_utils.ParseError(saveerr)

	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(err)

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
	_, saveerr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if saveerr != nil {
		return mysql_utils.ParseError(saveerr)

	}

	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, delErr := stmt.Exec(user.Id)

	if delErr != nil {
		return mysql_utils.ParseError(delErr)

	}

	return nil

}

func (user *User) FindByStatus(usr_status string) ([]User, *errors.RestErr) {
	stmt, err := user_db.Client.Prepare(querySearchUser)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	rows, search_err := stmt.Query(usr_status)
	if search_err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	res := make([]User, 0)
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.DateCreated, &usr.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		res = append(res, usr)
	}

	if len(res) == 0 {
		return nil, errors.NewNotFoundError("no users matching found")
	}

	return res, nil
}
