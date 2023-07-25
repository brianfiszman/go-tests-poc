package mocks_test

import "net/mail"

func getMailError(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
