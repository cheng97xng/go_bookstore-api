package users

import (
	"fmt"
	"go_bookstore-api/datasource/mysql/users_db"
	"go_bookstore-api/utils/date_utils"
	"go_bookstore-api/utils/errors"
	"go_bookstore-api/utils/mysql_utils"
)

//	func Get(userId int64) (*User, *errors.RestErr) {
//		return nil, nil
//	}
const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created)VALUES(?,?,?,?);"
	queryGetUser     = "SELECT * from users WHERE id=?;"
	queryUpdateUser  = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser  = "DELETE FROM users WHERE id=?;"
)

// var (
// 	usersDB = make(map[int64]*User)
// )

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		fmt.Println(getErr)
		// if strings.Contains(err.Error(), errorNoRows) {
		// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		fmt.Println(saveErr)
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		// return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
