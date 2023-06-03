package structs

import (
	"github.com/fatih/structs"
	"github.com/goadify/goadify/modules/crud/helpers"
	"github.com/goadify/goadify/modules/crud/models"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

const (
	jsonTagName                 = "json"
	jsonNotIncludeTag           = "-"
	errDatatypeNotSupportedText = "datatype (%s) not supported"
	omitEmptyTag                = ",omitempty"
)

func GetJsonFieldName(field *structs.Field) string {
	fieldName := field.Name()

	if !field.IsExported() {
		return ""
	}

	if field.IsEmbedded() {
		return ""
	}

	if jsonTag := field.Tag(jsonTagName); len(jsonTag) > 0 {
		fieldName = jsonTag

		if strings.Contains(jsonTag, omitEmptyTag) {
			oeTagPos := strings.Index(jsonTag, omitEmptyTag)
			jsonTag = jsonTagName[:oeTagPos]
		}

		if jsonTag == jsonNotIncludeTag {
			return ""
		}
	}

	return fieldName
}

func GetDatatype(value any) (dt models.Datatype, err error) {
	if helpers.IsInteger(value) {
		dt = models.DatatypeInteger
	} else if helpers.IsString(value) {
		dt = models.DatatypeString
	} else if helpers.IsFloat(value) {
		dt = models.DatatypeFloat
	} else {
		err = errors.Errorf(errDatatypeNotSupportedText, reflect.TypeOf(value).String())
	}

	return
}

func EntityToEntityMapping(entity any, name string) (entityMapping models.EntityMapping, errs []error) {
	entityMapping.Name = name

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

		f := models.Field{
			Name:     name,
			Datatype: datatype,
		}

		entityMapping.Fields = append(entityMapping.Fields, f)
	}

	return entityMapping, errs
}
