package qdoc

import (
	"github.com/getkin/kin-openapi/openapi3"
	"reflect"
)

func toOpenAPIType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.String:
		return "string"
	case reflect.Slice:
		return "array"
	case reflect.Map:
		return "object"
	case reflect.Struct:
		return "object"
	default:
		return "string"
	}
}

func toOpenAPISchema(t reflect.Type) *openapi3.Schema {
	if t == nil {
		return openapi3.NewSchema()
	}
	switch t.Kind() {
	case reflect.Bool:
		return &openapi3.Schema{
			Type: "boolean",
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &openapi3.Schema{
			Type: "integer",
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &openapi3.Schema{
			Type:   "integer",
			Format: "int64",
		}
	case reflect.Float32, reflect.Float64:
		return &openapi3.Schema{
			Type: "number",
		}
	case reflect.String:
		return &openapi3.Schema{
			Type: "string",
		}
	case reflect.Struct:
		properties := make(map[string]*openapi3.SchemaRef)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			properties[generateName(field)] = &openapi3.SchemaRef{
				Value: toOpenAPISchema(field.Type),
			}
		}
		return &openapi3.Schema{
			Type:       "object",
			Properties: properties,
		}
	case reflect.Slice:
		return &openapi3.Schema{
			Type: "array",
			Items: &openapi3.SchemaRef{
				Value: toOpenAPISchema(t.Elem()),
			},
		}
	default:
		panic("unsupported type")
	}
}

func generateName(t reflect.StructField) string {
	if tag := t.Tag.Get("json"); tag != "" {
		return tag
	}
	return t.Name
}
