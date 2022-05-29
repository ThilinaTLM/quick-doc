package qdoc

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
	Schema       SchemaConfig
	Description  string
}

type RespSet struct {
	Success   *Response
	BadReq    *Response
	UnAuth    *Response
	Forbidden *Response
	NotFound  *Response
	ISE       *Response
	others    map[HttpStatus]*Response
}

func (r *RespSet) collectToMap() map[HttpStatus]*Response {
	var m = make(map[HttpStatus]*Response)
	if r.Success != nil {
		m[HTTP_OK] = r.Success
	}
	if r.BadReq != nil {
		m[HTTP_BAD_REQUEST] = r.BadReq
	}
	if r.UnAuth != nil {
		m[HTTP_UNAUTHORIZED] = r.UnAuth
	}
	if r.Forbidden != nil {
		m[HTTP_FORBIDDEN] = r.Forbidden
	}
	if r.NotFound != nil {
		m[HTTP_NOT_FOUND] = r.NotFound
	}
	if r.ISE != nil {
		m[HTTP_ISE] = r.ISE
	}
	for k, v := range r.others {
		m[k] = v
	}

	// set status codes
	for k, v := range m {
		v.Status = k
	}

	// remove invalids
	for k, v := range m {
		if v.Status == 0 || v.ContentTypes == nil || len(v.ContentTypes) == 0 {
			delete(m, k)
		}
	}
	return m
}

// ResJson returns a Response with a json content type
func ResJson(desc string, value interface{}) *Response {
	return &Response{
		ContentTypes: []ContentType{CONTENT_TYPE_JSON},
		Schema:       Schema(value),
		Description:  desc,
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

func (r RespSet) toOpenAPI() openapi3.Responses {
	_responses := make(openapi3.Responses)
	for _, resp := range r.collectToMap() {
		_responses[strconv.Itoa(int(resp.Status))] = &openapi3.ResponseRef{
			Value: resp.toOpenAPI(),
		}
	}
	return _responses
}
