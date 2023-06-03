package crud

import (
	"context"
)

// Entity describes an entity and its identifier
type Entity interface {
	EntityName() string
	ID() any
}

// Record used for receive/retrieve data
type Record interface {
	ID() any
}

// Repository base interface for repositories.
// see RepositoryCreatable, RepositoryReadable, RepositoryUpdatable, RepositoryDeletable
type Repository interface {
	AccessRules(context.Context) ([]AccessRule, error)
}

type RepositoryCreatable interface {
	Create(context.Context, Record) error
}

type RepositoryReadable interface {
	GetByID(context.Context) (Record, error)
	GetList(ctx context.Context, page int32, perPage int32) (records []Record, totalCount int64, err error)
}

type RepositoryUpdatable interface {
	Update(context.Context, Record) error
}

type RepositoryDeletable interface {
	Delete(context.Context, Record) error
}
