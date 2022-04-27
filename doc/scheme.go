package doc

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
	"reflect"
	"strconv"
)

// Scheme is document data scheme configuration
func Scheme(value interface{}) SchemeConfig {
	return SchemeConfig{
		Object: value,
	}
}

type SchemeConfig struct {
	Object interface{}
}

func (sb *SchemeConfig) toOpenAPI() *openapi3.Schema {
	if sb.Object == nil {
		return openapi3.NewSchema()
	}

	_type := reflect.TypeOf(sb.Object)
	_value := reflect.ValueOf(sb.Object)

	scheme := convertToSchema(_type, _value)
	scheme.Example = filterExampleObject(sb.Object, scheme)
	return scheme
}

func convertToSchema(t reflect.Type, v reflect.Value) *openapi3.Schema {
	switch t.Kind() {
	case reflect.Interface:
		if v.Interface() == nil {
			return nil
		}
		return convertToSchema(v.Elem().Type(), v.Elem())
	case reflect.Ptr:
		if !v.IsValid() {
			return convertToSchema(t.Elem(), reflect.Value{})
		}
		return convertToSchema(t.Elem(), v.Elem())
	case reflect.Invalid:
		return openapi3.NewSchema()
	case reflect.Bool:
		return &openapi3.Schema{
			Type:    "boolean",
			Example: getExample(t, v),
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &openapi3.Schema{
			Type:    "integer",
			Example: getExample(t, v),
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &openapi3.Schema{
			Type:    "integer",
			Format:  "int64",
			Example: getExample(t, v),
		}
	case reflect.Float32, reflect.Float64:
		return &openapi3.Schema{
			Type:    "number",
			Example: getExample(t, v),
		}
	case reflect.String:
		return &openapi3.Schema{
			Type:    "string",
			Example: getExample(t, v),
		}
	case reflect.Struct:
		properties := make(map[string]*openapi3.SchemaRef)
		for i := 0; i < t.NumField(); i++ {
			_field := t.Field(i)
			var _value reflect.Value
			if v.IsValid() {
				_value = v.Field(i)
			} else {
				_value = reflect.Value{}
			}
			if _field.Type.Kind() == reflect.Ptr && _value.IsNil() {
				continue
			}
			_schema := convertToSchema(_field.Type, _value)
			if _schema == nil {
				continue
			}
			properties[jsonName(_field)] = &openapi3.SchemaRef{
				Value: _schema,
			}
		}
		return &openapi3.Schema{
			Type:       "object",
			Properties: properties,
		}
	case reflect.Slice:
		// TODO: support slice example v
		return &openapi3.Schema{
			Type: "array",
			Items: &openapi3.SchemaRef{
				Value: convertToSchema(t.Elem(), reflect.Value{}),
			},
			Example: v,
		}
	case reflect.Map:
		properties := make(map[string]*openapi3.SchemaRef)
		for _, key := range v.MapKeys() {
			sKey, ok := isStringOrNumberKind(key)
			if !ok {
				panic("unsupported type as map key, only string or number is required")
			}
			val := v.MapIndex(key)
			properties[sKey] = &openapi3.SchemaRef{
				Value: convertToSchema(val.Type(), val),
			}
		}
		return &openapi3.Schema{
			Type:       "object",
			Properties: properties,
		}

	default:
		panic("unsupported type")
	}
}

func getExample(t reflect.Type, v reflect.Value) interface{} {
	if v.Kind() == reflect.Invalid {
		return nil
	}

	switch t.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.String:
		return v.String()
	}

	return nil
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

func filterExampleObject(v interface{}, schema *openapi3.Schema) *map[string]interface{} {
	if v == nil {
		return nil
	}

	// convert to map
	_json, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var _map map[string]interface{}
	err = json.Unmarshal([]byte(_json), &_map)

	// remove null fields
	for k, v := range _map {
		if v == nil {
			delete(_map, k)
		}
	}
	return &_map
}
