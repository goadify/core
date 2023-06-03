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
// for extend functions see RepositoryCreatable, RepositoryReadable, RepositoryUpdatable, RepositoryDeletable
type Repository interface {
	// AccessRules Must return a set of access rules, that allowed for user
	AccessRules(context.Context) ([]AccessRule, error)

	// NewRecord Must return a new instance of model
	NewRecord() Record
}

type RepositoryCreatable interface {
	Create(context.Context, Record) error
}

type RepositoryReadable interface {
	GetByID(context.Context, string) (Record, error)
	GetList(ctx context.Context, page int32, perPage int32) (records []Record, totalCount int64, err error)
}

type RepositoryUpdatable interface {
	Update(context.Context, Record, string) error
}

type RepositoryDeletable interface {
	Delete(context.Context, string) error
}
