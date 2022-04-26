package doc

import (
	"github.com/getkin/kin-openapi/openapi3"
	"strconv"
)

type HttpStatus int

var HTTP_OK = HttpStatus(200)
var HTTP_BAD_REQUEST = HttpStatus(400)
var HTTP_UNAUTHORIZED = HttpStatus(401)
var HTTP_FORBIDDEN = HttpStatus(403)
var HTTP_NOT_FOUND = HttpStatus(404)
var HTTP_ISE = HttpStatus(500)

type Response struct {
	Status       HttpStatus
	ContentTypes []ContentType
	Schema       SchemeConfig
	Description  string
}

type Responses []Response

func Resp(responses ...Response) Responses {
	_resp := make(Responses, 0)
	for _, resp := range responses {
		if resp.Status == 0 || len(resp.ContentTypes) == 0 {
			continue
		}
		_resp = append(_resp, resp)
	}
	return _resp
}

// RespSuccess returns a Response with a 200 status code and the given media types.
func RespSuccess(description string, value interface{}) Response {
	return Response{
		Status:       HTTP_OK,
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Description:  description,
	}
}

// RespBadReq returns a Response with a 400 status code and the given media types.
func RespBadReq(description string, value interface{}) Response {
	return Response{
		Status:       HTTP_BAD_REQUEST,
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Description:  description,
	}
}

// RespUnauthorized returns a Response with a 401 status code and the given media types.
func RespUnauthorized(description string, value interface{}) Response {
	return Response{
		Status:       HTTP_UNAUTHORIZED,
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Description:  description,
	}
}

// RespNotFound returns a Response with a 404 status code and the given media types.
func RespNotFound(description string, value interface{}) Response {
	return Response{
		Status:       HTTP_NOT_FOUND,
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Description:  description,
	}
}

// RespISE returns a Response with a 500 status code and the given media types.
func RespISE(description string, value interface{}) Response {
	return Response{
		Status:       HTTP_ISE,
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Description:  description,
	}
}

// RespError returns a Response with a given status code and the given media types.
func RespError(status HttpStatus, description string, value interface{}) Response {
	return Response{
		Status:       status,
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Scheme(value),
		Description:  description,
	}
}

func (r Response) toOpenAPI() *openapi3.Response {
	consumes := make([]string, 0)
	for _, ct := range r.ContentTypes {
		consumes = append(consumes, string(ct))
	}
	return &openapi3.Response{
		Description: &r.Description,
		Content: openapi3.NewContentWithSchemaRef(
			openapi3.NewSchemaRef("", r.Schema.toOpenAPI()),
			consumes,
		),
	}
}

func (r Responses) toOpenAPI() openapi3.Responses {
	_responses := make(openapi3.Responses)
	for _, resp := range r {
		_responses[strconv.Itoa(int(resp.Status))] = &openapi3.ResponseRef{
			Value: resp.toOpenAPI(),
		}
	}
	return _responses
}
