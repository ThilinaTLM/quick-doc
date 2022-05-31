package qdoc

import "github.com/getkin/kin-openapi/openapi3"

type RequestBody struct {
	ContentTypes []ContentType
	Schema       *SchemaConfig
	Required     bool
}

func ReqBody(sc *SchemaConfig) func(...ContentType) RequestBody {
	return func(contentTypes ...ContentType) RequestBody {
		return RequestBody{
			ContentTypes: contentTypes,
			Schema:       sc,
			Required:     true,
		}
	}
}

func ReqJson(sc *SchemaConfig) RequestBody {
	return RequestBody{
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       sc,
		Required:     true,
	}
}

func ReqForm(sc *SchemaConfig) RequestBody {
	return RequestBody{
		ContentTypes: []ContentType{CONTENT_TYPE_FORM},
		Schema:       sc,
		Required:     true,
	}
}

func (rb *RequestBody) toOpenAPI() *openapi3.RequestBody {
	if len(rb.ContentTypes) == 0 {
		return &openapi3.RequestBody{
			Description: "",
			Required:    false,
			Content:     openapi3.NewContent(),
		}
	}
	consumes := make([]string, 0)
	for _, ct := range rb.ContentTypes {
		consumes = append(consumes, string(ct))
	}
	return &openapi3.RequestBody{
		Content: openapi3.NewContentWithSchemaRef(
			openapi3.NewSchemaRef("", rb.Schema.toOpenAPI()),
			consumes,
		),
		Required: rb.Required,
	}
}
