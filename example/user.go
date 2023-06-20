package main

import (
	"context"
	"github.com/goadify/goadify/modules/crud"
)

type User struct {
	ID   int32
	Name string
}

func (u *User) Identifier() any {
	return u.ID
}

type UserRepository struct {
}

func (u *UserRepository) AccessRules(_ context.Context) ([]crud.AccessRule, error) {
	return []crud.AccessRule{
		crud.AccessCreateRule,
		crud.AccessReadRule,
		crud.AccessUpdateRule,
		crud.AccessDeleteRule,
	}, nil
}

func (u *UserRepository) NewModel() crud.Model {
	return new(User)
}
