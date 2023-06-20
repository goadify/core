package hydrator

import (
	"encoding/json"
	"github.com/goadify/goadify/modules/crud/models"
	crudGen "github.com/goadify/openapi/crud/go/gen"
)

func Model(model models.IdentifiedModel) (crudGen.IdentifiedRecord, error) {
	recordData, err := json.Marshal(model.Data)
	if err != nil {
		return crudGen.IdentifiedRecord{}, err
	}

	return crudGen.IdentifiedRecord{
		ID:   model.ID,
		Data: recordData,
	}, nil
}

func Models(models []models.IdentifiedModel) ([]crudGen.IdentifiedRecord, error) {
	res := make([]crudGen.IdentifiedRecord, len(models))

	for i := 0; i < len(models); i++ {
		model, err := Model(models[i])
		if err != nil {
			return nil, err
		}

		res[i] = model
	}

	return res, nil
}
