package users

import (
	"fmt"
	"go_bookstore-api/datasource/mysql/users_db"
	"go_bookstore-api/utils/date_utils"
	"go_bookstore-api/utils/errors"
	"strings"
)

//	func Get(userId int64) (*User, *errors.RestErr) {
//		return nil, nil
//	}
const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created)VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

// func something() {
// 	user := User{}
// 	if err := user.Get(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(user.FirstName)
// }

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	// current := usersDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }

	// user.DateCreated = date_utils.GetNowString()

	// usersDB[user.Id] = user
	return nil
}
