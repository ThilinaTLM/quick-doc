package doc

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"reflect"
	"strconv"
)

// Scheme is document data scheme configuration
func Scheme(value interface{}) SchemeConfig {
	t := reflect.TypeOf(value)
	return SchemeConfig{
		Type:  t,
		Value: value,
	}
}

// SchemeWithTitle is document data scheme configuration
func SchemeWithTitle(value interface{}, title string) SchemeConfig {
	t := reflect.TypeOf(value)
	return SchemeConfig{
		Title: title,
		Type:  t,
		Value: value,
	}
}

type SchemeConfig struct {
	Title string
	Type  reflect.Type
	Value interface{}
}

func (sb *SchemeConfig) toOpenAPI() *openapi3.Schema {
	var vo reflect.Value
	if sb.Value != nil {
		vo = reflect.ValueOf(sb.Value)
	}
	scheme := typeToScheme(sb.Type, vo)
	scheme.Example = sb.Value
	scheme.Title = sb.Title
	return scheme
}

func getExampleValue(v interface{}) {

}

func typeToScheme(t reflect.Type, value reflect.Value) *openapi3.Schema {
	if t == nil {
		return openapi3.NewSchema()
	}
	switch t.Kind() {
	case reflect.Bool:
		return &openapi3.Schema{
			Type:    "boolean",
			Example: value.Bool(),
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &openapi3.Schema{
			Type:    "integer",
			Example: value.Int(),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &openapi3.Schema{
			Type:    "integer",
			Format:  "int64",
			Example: value.Uint(),
		}
	case reflect.Float32, reflect.Float64:
		return &openapi3.Schema{
			Type:    "number",
			Example: value.Float(),
		}
	case reflect.String:
		return &openapi3.Schema{
			Type:    "string",
			Example: value.String(),
		}
	case reflect.Struct:
		properties := make(map[string]*openapi3.SchemaRef)
		for i := 0; i < value.NumField(); i++ {
			_field := t.Field(i)
			_value := value.Type().Field(i)
			fmt.Println(_value)
			properties[jsonName(_field)] = &openapi3.SchemaRef{
				Value: typeToScheme(_field.Type, reflect.Value{}),
			}
		}
		return &openapi3.Schema{
			Type:       "object",
			Properties: properties,
		}
	case reflect.Slice:
		// TODO: support slice example value
		return &openapi3.Schema{
			Type: "array",
			Items: &openapi3.SchemaRef{
				Value: typeToScheme(t.Elem(), reflect.Value{}),
			},
			Example: value,
		}
	case reflect.Map:
		properties := make(map[string]*openapi3.SchemaRef)
		for _, key := range value.MapKeys() {
			sKey, ok := isStringOrNumberKind(key)
			if !ok {
				panic("unsupported type as map key, only string or number is required")
			}
			val := value.MapIndex(key)
			properties[sKey] = &openapi3.SchemaRef{
				Value: typeToScheme(val.Type(), val),
			}
		}
		return &openapi3.Schema{
			Type:       "object",
			Properties: properties,
		}
	case reflect.Interface:
		return typeToScheme(value.Elem().Type(), value.Elem())
	default:
		panic("unsupported type")
	}
}

func isStringOrNumberKind(k reflect.Value) (string, bool) {
	kind := k.Kind()
	if kind == reflect.String {
		return k.String(), true
	}
	if kind == reflect.Int || kind == reflect.Int8 ||
		kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
		return strconv.FormatInt(k.Int(), 10), true
	}
	return "", false
}

func jsonName(t reflect.StructField) string {
	if tag := t.Tag.Get("json"); tag != "" {
		return tag
	}
	return t.Name
}
