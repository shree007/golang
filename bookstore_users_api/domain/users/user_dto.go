package users

import (
	"bookstore_users_api/utils/errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	DateCreated string
}

func (user *User) Validate() *errors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if strings.TrimSpace(user.Email) == "" {
		return errors.NewBadReuestError("invalid email address")
	}
	fmt.Println(user.Email)
	FormatErr := checkmail.ValidateFormat(user.Email)
	if FormatErr != nil {
		return errors.NewBadReuestError("Invalid Format")
	}

	// Validate

	user.Phone = strings.TrimSpace(user.Phone)
	if user.Phone == "" {
		return errors.NewBadReuestError("phone number required")
	}
	if len(user.Phone) != 10 {
		return errors.NewBadReuestError("phone number length should be 10")
	}

	return nil
}
