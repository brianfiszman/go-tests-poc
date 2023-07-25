package interfaces

import "tddservice.com/pkg/domain/entities"

type UserBuilder interface {
	Builder[entities.User]
	WithName(name string) UserBuilder
	WithEmail(email string) UserBuilder
}
