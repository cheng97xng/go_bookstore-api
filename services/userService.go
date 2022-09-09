package services

import "go_bookstore-api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
