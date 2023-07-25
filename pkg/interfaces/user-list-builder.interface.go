package interfaces

import "tddservice.com/pkg/domain/entities"

type UserListBuilder interface {
	WithCustomFunc(customFunc func(UserBuilder)) UserListBuilder
	Build() []*entities.User
}
