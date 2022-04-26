package doc

import "github.com/getkin/kin-openapi/openapi3"

type RequestBody struct {
	ContentTypes []ContentType
	Schema       SchemeConfig
	Required     bool
}

func ReqBody(value interface{}) func(...ContentType) RequestBody {
	return func(contentTypes ...ContentType) RequestBody {
		return RequestBody{
			ContentTypes: contentTypes,
			Schema:       Scheme(value),
			Required:     true,
		}
	}
}

func ReqBodyJson(value interface{}) RequestBody {
	return RequestBody{
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Required:     true,
	}
}

func ReqBodyForm(value interface{}) RequestBody {
	return RequestBody{
		ContentTypes: []ContentType{CONTENT_TYPE_FORM},
		Schema:       Scheme(value),
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
