package users

import (
	"bookstore_users_api/datasources/msql/usersDb"
	"bookstore_users_api/utils/errors"
	"fmt"
)

var (
	userDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, phone) VALUES(?, ?, ?, ?, ?);"
)

func (user *User) Get() *errors.RestErr {

	if err := usersDb.Client.Ping(); err != nil {

		panic(err)
	}

	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.Phone = result.Phone
	user.DateCreated = result.DateCreated
	return nil

}

func (user *User) Save() *errors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryInsertUser)
	if err != nil {
		panic(err.Error())
	}
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Phone)
	if err != nil {
		return errors.NewInternalError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		errors.NewInternalError(fmt.Sprintf("Error when trying to get the last inserted id: %s", err.Error()))
	}
	user.Id = userId
	defer stmt.Close()
	return nil

	// current := userDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadReuestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadReuestError(fmt.Sprintf("user id %d already exists", user.Id))
	// }
	// user.DateCreated = time.Now().Format("01-02-2006 15:04:05")
	// userDB[user.Id] = user
}
