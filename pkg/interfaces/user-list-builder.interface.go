package interfaces

import "tddservice.com/pkg/domain/entities"

type UserListBuilder interface {
	ListBuilder[entities.User]
}
