package qdoc

import "github.com/getkin/kin-openapi/openapi3"

// AuthType Authentication types
type AuthType string

const (
	AUTH_TYPE_BASIC  = AuthType("basic")
	AUTH_TYPE_BEARER = AuthType("bearerAuth")
)

type AuthConf []AuthType

func NewAuthConf() *AuthConf {
	return &AuthConf{}
}

func (a *AuthConf) Contains(authType AuthType) bool {
	for _, t := range *a {
		if t == authType {
			return true
		}
	}
	return false
}

func (a *AuthConf) With(authType AuthType) {
	if !a.Contains(authType) {
		*a = append(*a, authType)
	}
}

func (a *AuthConf) WithBearer() *AuthConf {
	a.With(AUTH_TYPE_BEARER)
	return a
}

func (a *AuthConf) WithBasic() *AuthConf {
	a.With(AUTH_TYPE_BASIC)
	return a
}

func (a AuthType) toOpenAPI() *openapi3.SecurityScheme {
	switch a {
	case AUTH_TYPE_BASIC:
		return &openapi3.SecurityScheme{
			Type: "basic",
		}
	case AUTH_TYPE_BEARER:
		return &openapi3.SecurityScheme{
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		}
	}
	return nil
}

func (a AuthType) toOpenAPISecurityRequirement() openapi3.SecurityRequirement {
	return openapi3.NewSecurityRequirement().Authenticate(string(a))
}

func (e *Endpoint) WithAuth(authType AuthType) *Endpoint {
	e.auth = true
	e.authConf.With(authType)
	return e
}

func (e *Endpoint) WithBearerAuth() *Endpoint {
	e.auth = true
	e.authConf.WithBearer()
	return e
}

func (e *Endpoint) WithBasicAuth() *Endpoint {
	e.auth = true
	e.authConf.WithBasic()
	return e
}
