package repository

import (
	"context"
	"github.com/goadify/goadify/example/models"
	"github.com/goadify/goadify/modules/crud"
)

type UserRepository struct {
}

func (u *UserRepository) AccessRules(ctx context.Context) ([]crud.AccessRule, error) {
	return []crud.AccessRule{
		crud.AccessCreateRule,
		crud.AccessReadRule,
		crud.AccessUpdateRule,
		crud.AccessDeleteRule,
	}, nil
}

func (u *UserRepository) NewRecord() crud.Record {
	return new(models.User)
}

func (u *UserRepository) Create(ctx context.Context, record crud.Record) error {
	rec, ok := record.(*models.User)
	if !ok {
		return nil
	}

	rec.Id = 2

	return nil
}

func (u *UserRepository) GetByID(_ context.Context, _ string) (crud.Record, error) {
	return &models.User{
		Id:   1,
		Name: "Obema",
	}, nil
}

func (u *UserRepository) GetList(ctx context.Context, page int32, perPage int32) (records []crud.Record, totalCount int64, err error) {
	return []crud.Record{
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

func (u *UserRepository) Update(ctx context.Context, record crud.Record, s string) error {
	return nil
}

func (u *UserRepository) Delete(ctx context.Context, s string) error {
	return nil
}
