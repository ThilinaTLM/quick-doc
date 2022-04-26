package doc

import "github.com/getkin/kin-openapi/openapi3"

// AuthType Authentication types
type AuthType string

const (
	AUTH_TYPE_BASIC  = AuthType("basic")
	AUTH_TYPE_BEARER = AuthType("bearerAuth")
)

// AuthTypesBearer returns the authentication type required for the bearer token
func AuthTypesBearer() []AuthType {
	return []AuthType{AUTH_TYPE_BEARER}
}

// AuthTypesBasic returns the authentication type required for the basic token
func AuthTypesBasic() []AuthType {
	return []AuthType{AUTH_TYPE_BASIC}
}

// AuthTypes helper function for create auth types
func AuthTypes(authTypes ...AuthType) []AuthType {
	return authTypes
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
