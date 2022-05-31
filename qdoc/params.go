package qdoc

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type ParamType string

const (
	PARAM_TYPE_QUERY  = ParamType("query")
	PARAM_TYPE_PATH   = ParamType("path")
	PARAM_TYPE_HEADER = ParamType("header")
)

type Parameter struct {
	Name        string
	Scheme      *SchemaConfig
	Description string
	Required    bool
	Loc         ParamType
}

type Parameters []Parameter

func PathParams(params ...Parameter) Parameters {
	for i, param := range params {
		param.Loc = PARAM_TYPE_PATH
		params[i] = param
	}
	return params
}

func Headers(headers ...Parameter) Parameters {
	for i, header := range headers {
		header.Loc = PARAM_TYPE_HEADER
		headers[i] = header
	}
	return headers
}

func QueryParams(params ...Parameter) Parameters {
	for i, param := range params {
		param.Loc = PARAM_TYPE_QUERY
		params[i] = param
	}
	return params
}

// RequiredParam returns a Parameter with the given name and value,
func RequiredParam(name string, value *SchemaConfig) Parameter {
	return Parameter{
		Name:     name,
		Scheme:   value,
		Required: true,
	}
}

// OptionalParam returns a Parameter with the given name and value,
func OptionalParam(name string, value *SchemaConfig) Parameter {
	return Parameter{
		Name:     name,
		Scheme:   value,
		Required: false,
	}
}

func (p *Parameter) toOpenAPI() *openapi3.Parameter {
	return &openapi3.Parameter{
		Name:        p.Name,
		In:          string(p.Loc),
		Description: p.Description,
		Required:    p.Required,
		Schema:      openapi3.NewSchemaRef("", p.Scheme.toOpenAPI()),
	}
}
