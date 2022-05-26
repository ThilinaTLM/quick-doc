package schema

import (
	"fmt"
	"reflect"
)

type Options struct {
	FollowPointers bool
}

type Builder struct {
	Options *Options
}

func FromObject(obj interface{}) []Property {
	if obj == nil {
		return []Property{}
	}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	properties, err := traverseStruct(t, v)
	if err != nil {
		panic(err)
	}
	return properties
}

func traverseStruct(t reflect.Type, v reflect.Value) ([]Property, error) {
	properties := make([]Property, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			continue
		}
		property := Property{}
		property.Name = getStructFieldName(field)
		property.Type = getPropType(field.Type)
		property.Description = getStructFieldDesc(field)
		property.Value = fmt.Sprintf("%v", v.Field(i).Interface())
		properties = append(properties, property)
	}
	return properties, nil
}
