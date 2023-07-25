package entities

import (
	"errors"
	"net/mail"
)

const (
	ERR_NAME_SHORT = "username is too short"
	ERR_NAME_EMPTY = "username is empty"
)

type User struct {
	Username string `fake:"{username}"`
	Email    string `fake:"{email}"`
}

func NewUser(username string, email string) (u *User) {
	u = &User{username, email}

	return
}

func (u *User) IsUserValid() error {
	if err := u.isUsernameValid(); err != nil {
		return err
	}

	if err := u.isEmailValid(); err != nil {
		return err
	}

	return nil
}

func (u *User) isUsernameValid() error {
	if err := u.isUsernameEmpty(); err != nil {
		return err
	}

	if err := u.isUsernameShort(); err != nil {
		return err
	}

	return nil
}

func (u *User) isUsernameEmpty() error {
	if err := errors.New(ERR_NAME_EMPTY); u.Username == "" {
		return err
	}

	return nil
}

func (u *User) isUsernameShort() error {
	if err := errors.New(ERR_NAME_SHORT); len(u.Username) < 6 {
		return err
	}

	return nil
}

func (u *User) isEmailValid() (err error) {
	_, err = mail.ParseAddress(u.Email)

	return err
}
