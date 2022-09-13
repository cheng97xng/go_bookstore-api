package services

import (
	"go_bookstore-api/domain/users"
	"go_bookstore-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// if err := users.Validate(&user); err != nil {
	// 	return nil, err
	// }

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	// current:=users.User{Id: user.Id}
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	//validate email
	if err := user.Validate(); err != nil {
		return nil, err
	}
	//check update email only
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
	return current, nil
}
