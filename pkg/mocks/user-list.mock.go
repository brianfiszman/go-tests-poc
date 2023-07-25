package mocks

import (
	"tddservice.com/pkg/domain/entities"
	"tddservice.com/pkg/interfaces"
)

type UserListBuilder struct {
	UserBuilderList []interfaces.UserBuilder
	quantity        int
}

func NewUserListBuilder(quantity int) interfaces.UserListBuilder {
	var builderList []interfaces.UserBuilder = make([]interfaces.UserBuilder, 0, quantity)

	for i := 0; i < quantity; i++ {
		builderList = append(builderList, NewUserBuilder())
	}

	return &UserListBuilder{quantity: quantity, UserBuilderList: builderList}
}

func (u *UserListBuilder) WithCustomFunc(customFunc func(interfaces.UserBuilder)) interfaces.UserListBuilder {
	for i := 0; i < u.quantity; i++ {
		customFunc(u.UserBuilderList[i])
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
