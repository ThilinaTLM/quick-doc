package qdoc

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
)

type CompiledDoc struct {
	Json   []byte
	config Config
	specs  *openapi3.T
}

func (d *Doc) Compile() (*CompiledDoc, error) {
	spec := d.compileSpecs(d)
	err := spec.Validate(context.Background())
	if err != nil {
		return nil, err
	}
	bytes, err := spec.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return &CompiledDoc{
		config: d.config,
		specs:  spec,
		Json:   bytes,
	}, nil
}

func (d *Doc) compileSpecs(doc *Doc) *openapi3.T {
	spec := openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       doc.config.Title,
			Description: doc.config.Description,
			Version:     doc.config.Version,
		},
		Servers: d.compileServerList(),
		Paths:   d.compilePaths(),
		Components: openapi3.Components{
			SecuritySchemes: d.compileSecuritySchemes(),
		},
	}
	return &spec
}

func (d *Doc) compileServerList() []*openapi3.Server {
	urls := d.config.Servers
	servers := make([]*openapi3.Server, len(urls))
	for i, s := range urls {
		servers[i] = &openapi3.Server{
			URL: s,
		}
	}
	return servers
}

func (d *Doc) compileSecuritySchemes() openapi3.SecuritySchemes {
	var types AuthConf

	for _, v := range *d.config.AuthConf {
		if !types.Contains(v) {
			types = append(types, v)
		}
	}

	for _, ep := range d.endpoints {
		if ep.auth {
			for _, v := range ep.authConf {
				if !types.Contains(v) {
					types = append(types, v)
				}
			}
		}
	}

	securitySchemes := make(openapi3.SecuritySchemes)
	for _, authType := range types {
		securitySchemes[string(authType)] = &openapi3.SecuritySchemeRef{
			Value: authType.toOpenAPI(),
		}
	}
	return securitySchemes
}

func (d *Doc) compilePaths() openapi3.Paths {
	paths := make(openapi3.Paths)
	for _, e := range d.endpoints {
		path, method, item := d.compileOperation(e)
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

func (d *Doc) compileOperation(ep *Endpoint) (path string, method MethodType, item openapi3.Operation) {
	path = ep.Path
	item = openapi3.Operation{
		Summary:     ep.Summary,
		Description: ep.Desc,
		RequestBody: &openapi3.RequestBodyRef{Value: ep.ReqBody.toOpenAPI()},
		Responses:   ep.RespSet.toOpenAPI(),
		Tags:        ep.tags,
		Parameters:  d.compileParams(ep.PathParams, ep.QueryParams, ep.Headers),
	}
	if ep.auth {
		item.Security = openapi3.NewSecurityRequirements()
		for _, authType := range ep.authConf {
			item.Security.With(authType.toOpenAPISecurityRequirement())
		}
	}
	method = ep.method
	return
}

// compileParams converts the parameters into a list of openapi3.Parameters
func (d *Doc) compileParams(paramSet ...Parameters) openapi3.Parameters {
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
