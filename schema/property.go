package schema

import (
	"fmt"
	"reflect"
)

// PropType Schema Property Type
type PropType string

const (
	PropType_INTEGER PropType = "INTEGER"
	PropType_NUMBER  PropType = "NUMBER"
	PropType_BOOLEAN PropType = "BOOLEAN"
	PropType_STRING  PropType = "STRING"
	PropType_OBJECT  PropType = "OBJECT"
	PropType_ARRAY   PropType = "ARRAY"
	PropType_MAP     PropType = "MAP"
)

// GetPropType returns corresponding PropType of the given reflect.Type
func GetPropType(t reflect.Type) PropType {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return PropType_INTEGER
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return PropType_INTEGER
	case reflect.Float32, reflect.Float64:
		return PropType_NUMBER
	case reflect.Bool:
		return PropType_BOOLEAN
	case reflect.String:
		return PropType_STRING
	case reflect.Struct:
		return PropType_OBJECT
	case reflect.Slice, reflect.Array:
		return PropType_ARRAY
	case reflect.Map:
		return PropType_MAP
	default:
		panic(fmt.Sprintf("unsupported type: %s", t.Kind()))
	}
}

type Property struct {
	Type        PropType     `json:"type"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Value       string       `json:"value"`
	Properties  []Property   `json:"properties"`
	Constraints []Constraint `json:"constraints"`
}

func (p *Property) WithName(s string) *Property {
	p.Name = s
	return p
}

func (p *Property) WithDesc(s string) *Property {
	p.Description = s
	return p
}

func (p *Property) WithConstraint(c Constraint) *Property {
	if p.Constraints == nil {
		p.Constraints = make([]Constraint, 0)
	}
	p.Constraints = append(p.Constraints, c)
	return p
}
