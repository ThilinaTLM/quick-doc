package schema

import (
	"fmt"
	"reflect"
	"strings"
)

func NewBuilder() Builder {
	return Builder{
		Options: &Options{
			ExploreNilStruct: true,
			PreferJsonTag:    true,
		},
	}
}

type Options struct {
	ExploreNilStruct bool
	PreferJsonTag    bool
}

type Builder struct {
	Options *Options
}

func (b *Builder) GetSchema(obj interface{}) (*Property, error) {
	if obj == nil {
		return nil, nil
	}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	return b.inspect(t, v)
}

func (b *Builder) inspect(t reflect.Type, v reflect.Value) (*Property, error) {
	switch t.Kind() {
	case reflect.Interface:
		if v.Interface() == nil {
			return nil, nil
		}
		return b.inspect(v.Elem().Type(), v.Elem())
	case reflect.Ptr:
		if !v.IsValid() {
			return b.inspect(t.Elem(), reflect.Value{})
		}
		return b.inspect(t.Elem(), v.Elem())
	case reflect.String:
		return &Property{
			Type:  PropType_STRING,
			Value: b.valueString(v),
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &Property{
			Type:  PropType_INTEGER,
			Value: b.valueString(v),
		}, nil
	case reflect.Bool:
		return &Property{
			Type:  PropType_BOOLEAN,
			Value: b.valueString(v),
		}, nil
	case reflect.Struct:
		props, err := b.inspectStruct(t, v)
		if err != nil {
			return nil, err
		}
		return &Property{
			Type:       PropType_OBJECT,
			Properties: props,
		}, nil
	case reflect.Slice:
		props := make([]Property, 0)
		if v.Len() == 0 {
			prop, err := b.inspect(t.Elem(), reflect.Value{})
			if err != nil {
				return nil, err
			}
			props = append(props, *prop)
		} else {
			for i := 0; i < v.Len(); i++ {
				prop, err := b.inspect(t.Elem(), v.Index(i))
				if err != nil {
					return nil, err
				}
				props = append(props, *prop)
			}
		}
		return &Property{
			Type:       PropType_ARRAY,
			Properties: props,
		}, nil
	case reflect.Map:
		panic("not implemented")
	default:
		panic("unknown type")
	}
}

func (b *Builder) inspectStruct(t reflect.Type, v reflect.Value) ([]Property, error) {
	props := make([]Property, 0)
	for i := 0; i < t.NumField(); i++ {
		_field := t.Field(i)
		var _value reflect.Value
		if v.IsValid() {
			_value = v.Field(i)
		} else {
			_value = reflect.Value{}
		}
		prop, err := b.inspect(_field.Type, _value)
		if err != nil {
			return nil, err
		}
		if prop == nil {
			continue
		}
		prop = prop.
			WithName(b.structFieldName(_field))
		props = append(props, *prop)
	}
	return props, nil
}

func (b *Builder) structFieldName(sf reflect.StructField) string {
	if b.Options.PreferJsonTag {
		jsonTag := strings.Split(sf.Tag.Get("json"), ",")[0]
		if jsonTag != "" {
			return jsonTag
		}
	}
	return sf.Name
}

func (b *Builder) valueString(v reflect.Value) string {
	if !v.IsValid() {
		return "nil"
	}
	return fmt.Sprintf("%v", v)
}
