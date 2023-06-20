package crud

import (
	"context"
	"github.com/goadify/goadify/modules/navigation"
)

type Model interface {
	Identifier() any
}

type Repository interface {
	// AccessRules Must return a set of access rules, that allowed for context
	AccessRules(context.Context) ([]AccessRule, error)

	// NewModel Must return a new instance of model
	NewModel() Model
}

type RepositoryCreatable interface {
	Create(context.Context, Model) error
}

type RepositoryReadable interface {
	GetByID(context.Context, string) (Model, error)
	GetList(ctx context.Context, page int32, perPage int32) (models []Model, totalCount int64, err error)
}

type RepositoryUpdatable interface {
	Update(context.Context, Model, string) error
}

type RepositoryDeletable interface {
	Delete(context.Context, string) error
}

type Link struct {
	Title    string
	Priority int64
}

type Group struct {
	Title    string
	Priority int64
}

type Entity struct {
	Slug string
	Name string

	Repository Repository

	Link  *navigation.Link
	Group *navigation.Group
}
