package interfaces

import (
	"tddservice.com/pkg/domain/entities"
)

type ListBuilder interface {
	WithCustomFunc(customFunc func(*Builder[any])) *ListBuilder
	Build() []*entities.User
}
