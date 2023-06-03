package structs

import (
	"github.com/fatih/structs"
	"github.com/goadify/goadify/modules/crud/helpers"
	models2 "github.com/goadify/goadify/modules/crud/models"
	"github.com/pkg/errors"
	"reflect"
)

const (
	JsonTagName                 = "json"
	JsonNotIncludeTag           = "-"
	ErrDatatypeNotSupportedText = "datatype (%s) not supported"
)

type Entity interface {
	EntityName() string
}

func GetJsonFieldName(field *structs.Field) string {
	fieldName := field.Name()

	if !field.IsExported() {
		return ""
	}

	if field.IsEmbedded() {
		return ""
	}

	if jsonTag := field.Tag(JsonTagName); len(jsonTag) > 0 {
		fieldName = jsonTag

		if jsonTag == JsonNotIncludeTag {
			return ""
		}
	}

	return fieldName
}

func GetDatatype(value any) (dt models2.Datatype, err error) {
	if helpers.IsInteger(value) {
		dt = models2.DatatypeInteger
	} else if helpers.IsString(value) {
		dt = models2.DatatypeString
	} else if helpers.IsFloat(value) {
		dt = models2.DatatypeFloat
	} else {
		err = errors.Errorf(ErrDatatypeNotSupportedText, reflect.TypeOf(value).String())
	}

	return dt, err
}

func EntityToEntityMapping(entity Entity) (entityMapping models2.EntityMapping, errs []error) {
	entityMapping.Name = entity.EntityName()

	strHelper := structs.New(entity)
	fields := strHelper.Fields()
	values := strHelper.Map()
	for _, field := range fields {

		name := GetJsonFieldName(field)
		if len(name) == 0 {
			continue
		}

		datatype, err := GetDatatype(values[field.Name()])
		if err != nil {
			errs = append(errs, err)
			continue
		}

		f := models2.Field{
			Name:     name,
			Datatype: datatype,
		}

		entityMapping.Fields = append(entityMapping.Fields, f)
	}

	return entityMapping, errs
}
