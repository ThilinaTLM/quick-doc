package doc

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
)

func (d *Doc) Compile() (*CompiledDoc, error) {
	spec := buildOpenAPISpec(d)
	err := spec.Validate(context.Background())
	if err != nil {
		return nil, err
	}
	bytes, err := spec.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return &CompiledDoc{
		raw:  d,
		doc:  spec,
		Json: bytes,
	}, nil
}

type CompiledDoc struct {
	raw  *Doc
	doc  *openapi3.T
	Json []byte
}

func buildOpenAPISpec(doc *Doc) *openapi3.T {
	spec := openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       doc.config.Title,
			Description: doc.config.Description,
			Version:     doc.config.Version,
		},
		Servers: buildOpenAPIServers(doc.config.Servers),
		Paths:   compileEndpoints(doc.endpoints),
		Components: openapi3.Components{
			SecuritySchemes: compileSecuritySchemes(doc.config.AuthTypes),
		},
	}
	return &spec
}

func buildOpenAPIServers(urls []string) []*openapi3.Server {
	servers := make([]*openapi3.Server, len(urls))
	for i, s := range urls {
		servers[i] = &openapi3.Server{
			URL: s,
		}
	}
	return servers
}

func compileSecuritySchemes(authTypes []AuthType) openapi3.SecuritySchemes {
	securitySchemes := make(openapi3.SecuritySchemes)
	for _, authType := range authTypes {
		securitySchemes[string(authType)] = &openapi3.SecuritySchemeRef{
			Value: authType.toOpenAPI(),
		}
	}
	return securitySchemes
}

func compileEndpoints(endpoints []Endpoint) openapi3.Paths {
	paths := make(openapi3.Paths)
	for _, e := range endpoints {
		path, method, item := compileOperation(e)
		pi := paths[path]
		if pi == nil {
			pi = &openapi3.PathItem{}
			paths[path] = pi
		}
		if method == METHOD_GET {
			pi.Get = &item
		} else if method == METHOD_POST {
			pi.Post = &item
		} else if method == METHOD_PUT {
			pi.Put = &item
		} else if method == METHOD_DELETE {
			pi.Delete = &item
		}
	}
	return paths
}

func compileOperation(ep Endpoint) (path string, method MethodType, item openapi3.Operation) {
	path = ep.Path
	item = openapi3.Operation{
		Summary:     ep.Summary,
		Description: ep.Description,
		RequestBody: &openapi3.RequestBodyRef{Value: ep.RequestBody.toOpenAPI()},
		Responses:   ep.Responses.toOpenAPI(),
		Tags:        ep.Tags,
		Parameters:  compileParams(ep.PathParams, ep.QueryParams, ep.Headers),
	}
	if len(ep.AuthTypes) > 0 {
		item.Security = openapi3.NewSecurityRequirements()
		for _, authType := range ep.AuthTypes {
			item.Security.With(authType.toOpenAPISecurityRequirement())
		}
	}
	method = ep.method
	return
}

// compileParams converts the parameters into a list of openapi3.Parameters
func compileParams(paramSet ...Parameters) openapi3.Parameters {
	_params := make(openapi3.Parameters, 0)
	for _, params := range paramSet {
		for _, p := range params {
			_params = append(_params, &openapi3.ParameterRef{
				Value: p.toOpenAPI(),
			})
		}
	}
	return _params
}
