package entities_test

import (
	"errors"
	"net/mail"
	"testing"

	"github.com/stretchr/testify/assert"
	"tddservice.com/pkg/domain/entities"
)

func getMailError(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func TestUser(t *testing.T) {
	t.Run("building default object", func(t *testing.T) {
		user := entities.NewUser("foobar", "foobar@gmail.com")

		assert.NotNil(t, user)
		assert.NoError(t, user.IsUserValid())
	})

	t.Run("building instances with faulty data", func(t *testing.T) {
		// Defining faulty users
		testCases := []struct {
			*entities.User
			expected error
		}{
			{User: entities.NewUser("foo", "foobar@gmail.com"), expected: errors.New(entities.ERR_NAME_SHORT)},
			{User: entities.NewUser("", "foobar@gmail.com"), expected: errors.New(entities.ERR_NAME_EMPTY)},
			{User: entities.NewUser("foobar", "foobar.gmail.com"), expected: getMailError("foobar.gmail.com")},
		}

		for _, tc := range testCases {
			t.Run(tc.expected.Error(), func(t *testing.T) {
				// Asserting Error with short name
				assert.NotNil(t, tc.User)
				assert.EqualError(t, tc.IsUserValid(), tc.expected.Error())
			})
		}
	})
}
