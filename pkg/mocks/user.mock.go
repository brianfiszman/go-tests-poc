package mocks

import (
	"github.com/brianvoe/gofakeit/v6"
	"tddservice.com/pkg/domain/entities"
)

type UserBuilder struct {
	user *entities.User
}

// NewUserBuilder creates an User Instance with default values
func NewUserBuilder() *UserBuilder {
	user := &entities.User{}

	gofakeit.Struct(user)

	return &UserBuilder{
		user: user,
	}
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Username = name

	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email

	return b
}

func (b *UserBuilder) Build() *entities.User {
	return b.user
}
