package interfaces

import "tddservice.com/pkg/domain/entities"

// UserBuilder interface defines the methods that a user builder should implement
type UserBuilder interface {
	WithName(name string) UserBuilder
	WithEmail(email string) UserBuilder
	Build() *entities.User
}
