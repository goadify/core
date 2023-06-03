package hydrator

import (
	"encoding/json"
	"github.com/goadify/goadify/modules/crud/models"
	crudGen "github.com/goadify/openapi/crud/go/gen"
)

func Record(record models.IdentifiedRecord) (crudGen.IdentifiedRecord, error) {
	recordData, err := json.Marshal(record.Data)
	if err != nil {
		return crudGen.IdentifiedRecord{}, err
	}

	return crudGen.IdentifiedRecord{
		ID:   record.ID,
		Data: recordData,
	}, nil
}

func Records(records []models.IdentifiedRecord) ([]crudGen.IdentifiedRecord, error) {
	res := make([]crudGen.IdentifiedRecord, len(records))

	for i := 0; i < len(records); i++ {
		model, err := Record(records[i])
		if err != nil {
			return nil, err
		}

		res[i] = model
	}

	return res, nil
}
