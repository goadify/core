package crud

import "context"

type Entity interface {
	EntityName() string
	ID() any
}

type Repository interface {
	AccessRules(context.Context) ([]AccessRule, error)
}

type RepositoryCreatable interface {
	Create(context.Context, Entity) error
}

type RepositoryReadable interface {
	GetByID(context.Context) (Entity, error)
	GetList(ctx context.Context, page int32, perPage int32) (entities []Entity, totalCount int32, err error)
}

type RepositoryUpdatable interface {
	Update(context.Context, Entity) error
}

type RepositoryDeletable interface {
	Delete(context.Context, Entity) error
}
