package repository

import (
	"context"
	"github.com/goadify/goadify/example/models"
	crudV1 "github.com/goadify/goadify/modules/crud"
)

type UserRepository struct {
}

func (u *UserRepository) GetByID(ctx context.Context) (crudV1.Record, error) {
	return &models.User{
		Id:   1,
		Name: "Obema",
	}, nil
}

func (u *UserRepository) GetList(ctx context.Context, page int32, perPage int32) (records []crudV1.Record, totalCount int64, err error) {
	return []crudV1.Record{
		&models.User{
			Id:   1,
			Name: "Obema",
		},
		&models.User{
			Id:   2,
			Name: "Trump",
		},
		&models.User{
			Id:   3,
			Name: "Putin",
		},
	}, 3, nil
}

func (u *UserRepository) AccessRules(ctx context.Context) ([]crudV1.AccessRule, error) {
	return []crudV1.AccessRule{
		crudV1.AccessCreateRule,
		crudV1.AccessReadRule,
		crudV1.AccessUpdateRule,
		crudV1.AccessDeleteRule,
	}, nil
}
