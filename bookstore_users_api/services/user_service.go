package services

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils/errors"
	"fmt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	fmt.Println(user)

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	fmt.Println(&user)
	return &user, nil

}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadReuestError("invalid user Id")
	}
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
