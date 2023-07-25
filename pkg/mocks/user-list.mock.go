package mocks

import (
	"tddservice.com/pkg/domain/entities"
)

type UserListBuilder struct {
	UserBuilderList []*UserBuilder
	quantity        int
}

func NewUserListBuilder(quantity int) *UserListBuilder {
	var builderList []*UserBuilder = make([]*UserBuilder, 0, quantity)

	for i := 0; i < quantity; i++ {
		builderList = append(builderList, NewUserBuilder())
	}

	return &UserListBuilder{quantity: quantity, UserBuilderList: builderList}
}

func (u *UserListBuilder) WithCustomFunc(customFunc func(*UserBuilder)) *UserListBuilder {
	for i := 0; i < u.quantity; i++ {
		customFunc(u.UserBuilderList[i])
		u.UserBuilderList = append(u.UserBuilderList, u.UserBuilderList[i])
	}

	return u
}

func (u *UserListBuilder) Build() []*entities.User {
	users := make([]*entities.User, 0, u.quantity)

	for _, builder := range u.UserBuilderList {
		users = append(users, builder.Build())
	}

	return users
}
