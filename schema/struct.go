package schema

import (
	"github.com/ThilinaTLM/quick-doc/helpers"
	"reflect"
)

func getStructFieldName(sf reflect.StructField) string {
	jsonTag := sf.Tag.Get("json")
	if jsonTag != "" {
		return jsonTag
	}
	return helpers.ToCamelCase(sf.Name, false)
}

func getStructFieldDesc(sf reflect.StructField) string {
	return sf.Tag.Get("qd")
}
