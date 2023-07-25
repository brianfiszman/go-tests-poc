package mocks_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"tddservice.com/pkg/domain/entities"
	"tddservice.com/pkg/interfaces"
	"tddservice.com/pkg/mocks"
)

func TestBuilder(t *testing.T) {
	var builder interfaces.UserBuilder = mocks.NewUserBuilder()

	t.Run("building a user using the builder", func(t *testing.T) {
		user := builder.Build()

		assert.NotNil(t, user)
		assert.NoError(t, user.IsUserValid())
	})

	t.Run("building instances with faulty data", func(t *testing.T) {
		// Defining faulty users
		testCases := []struct {
			interfaces.UserBuilder
			expected error
		}{
			{UserBuilder: mocks.NewUserBuilder().WithName("foo"), expected: errors.New(entities.ERR_NAME_SHORT)},
			{UserBuilder: mocks.NewUserBuilder().WithName(""), expected: errors.New(entities.ERR_NAME_EMPTY)},
			{UserBuilder: mocks.NewUserBuilder().WithEmail("foobar.gmail.com"), expected: getMailError("foobar.gmail.com")},
			{UserBuilder: mocks.NewUserBuilder().WithEmail(""), expected: getMailError("")},
		}

		for _, tc := range testCases {
			t.Run(tc.expected.Error(), func(t *testing.T) {
				// Build the user
				user := tc.Build()

				// Asserting Errors
				assert.NotNil(t, user)
				assert.EqualError(t, user.IsUserValid(), tc.expected.Error())
			})
		}
	})
}
