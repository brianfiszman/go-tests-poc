package mocks_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"tddservice.com/pkg/domain/entities"
	"tddservice.com/pkg/interfaces"
	"tddservice.com/pkg/mocks"
)

func TestBuilderList(t *testing.T) {
	t.Run("building a user using the builder", func(t *testing.T) {
		var testCases = []struct {
			interfaces.UserListBuilder
			expected error
		}{{
			UserListBuilder: mocks.NewUserListBuilder(100),
			expected:        nil,
		}}

		for _, tc := range testCases {
			t.Run("No Errors", func(t *testing.T) {
				users := tc.Build()

				assert.NotEmpty(t, users)
				assert.Len(t, users, 100)
				assert.IsType(t, []*entities.User{}, users)

				for _, u := range users {
					assert.NotNil(t, u)
					assert.NoError(t, u.IsUserValid())
				}
			})
		}
	})

	t.Run("building instances with faulty data", func(t *testing.T) {
		// Defining faulty users
		testCases := []struct {
			interfaces.UserListBuilder
			expected error
		}{
			{UserListBuilder: createUserListWithName(100, "foo"), expected: errors.New(entities.ERR_NAME_SHORT)},
			{UserListBuilder: createUserListWithName(100, ""), expected: errors.New(entities.ERR_NAME_EMPTY)},
			{UserListBuilder: createUserListWithEmail(100, "foobar.gmail.com"), expected: getMailError("foobar.gmail.com")},
			{UserListBuilder: createUserListWithEmail(100, ""), expected: getMailError("")},
		}

		for _, tc := range testCases {
			t.Run(tc.expected.Error(), func(t *testing.T) {
				// Build the user
				userList := tc.Build()

				assert.Len(t, userList, 100)

				for _, user := range userList {
					// Asserting Error with short name
					assert.NotNil(t, user)
					assert.EqualError(t, user.IsUserValid(), tc.expected.Error())
				}
			})
		}
	})
}

func createUserListWithName(quantity int, name string) interfaces.UserListBuilder {
	return mocks.NewUserListBuilder(quantity).WithCustomFunc(func(b interfaces.UserBuilder) { b.WithName(name) })
}

func createUserListWithEmail(quantity int, email string) interfaces.UserListBuilder {
	return mocks.NewUserListBuilder(quantity).WithCustomFunc(func(b interfaces.UserBuilder) { b.WithEmail(email) })
}
