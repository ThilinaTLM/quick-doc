package qdoc

import "github.com/ThilinaTLM/quick-doc/ui"

// MethodType Http methods
type MethodType string

const (
	METHOD_GET    = MethodType("GET")
	METHOD_POST   = MethodType("POST")
	METHOD_PUT    = MethodType("PUT")
	METHOD_DELETE = MethodType("DELETE")
)

type ContentType string

const (
	CONTENT_TYPE_JSON      = ContentType("application/json")
	CONTENT_TYPE_FORM      = ContentType("application/x-www-form-urlencoded")
	CONTENT_TYPE_MULTIPART = ContentType("multipart/form-data")
	CONTENT_TYPE_HTML      = ContentType("text/html")
	CONTENT_TYPE_FILE      = ContentType("application/octet-stream")
)

type UiConfig struct {
	Enabled      bool
	Path         string
	DefaultTheme ui.Theme
	ThemeByQuery bool
}

// Config API Doc configuration
type Config struct {
	Title       string
	Description string
	Version     string
	Servers     []string
	AuthConf    *AuthConf
	SpecPath    string

	UiConfig UiConfig
}

type Endpoint struct {
	Summary string
	Desc    string
	Path    string
	method  MethodType

	ReqBody     RequestBody
	QueryParams Parameters
	PathParams  Parameters
	Headers     Parameters
	RespSet     RespSet

	auth     bool
	authConf AuthConf
	tags     []string
}

type Doc struct {
	config    Config
	endpoints []*Endpoint
	schemas   []*SchemaConfig
}

func NewDoc(config Config) *Doc {

	if config.SpecPath == "" {
		config.SpecPath = "/doc/openapi.json"
	}

	if config.UiConfig.Enabled {
		if config.UiConfig.DefaultTheme == "" {
			config.UiConfig.DefaultTheme = ui.SWAGGER_UI
		}
		if config.UiConfig.Path == "" {
			config.UiConfig.Path = "/doc/ui"
		}
	}

	if config.AuthConf == nil {
		config.AuthConf = NewAuthConf()
	}

	return &Doc{
		config:    config,
		endpoints: make([]*Endpoint, 0),
	}
}

func (d *Doc) addEndpoint(ep *Endpoint) *Endpoint {
	ep.authConf = make([]AuthType, 0)
	ep.tags = make([]string, 0)
	d.endpoints = append(d.endpoints, ep)
	return ep
}

// Get add get endpoint
func (d *Doc) Get(ep *Endpoint) *Endpoint {
	ep.method = METHOD_GET
	return d.addEndpoint(ep)
}

// Post add post endpoint
func (d *Doc) Post(ep *Endpoint) *Endpoint {
	ep.method = METHOD_POST
	return d.addEndpoint(ep)
}

// Put add put endpoint
func (d *Doc) Put(ep *Endpoint) *Endpoint {
	ep.method = METHOD_PUT
	return d.addEndpoint(ep)
}

// Delete add delete endpoint
func (d *Doc) Delete(ep *Endpoint) *Endpoint {
	ep.method = METHOD_DELETE
	return d.addEndpoint(ep)
}

func (e *Endpoint) Tag(tag string) *Endpoint {
	e.tags = append(e.tags, tag)
	return e
}

func Servers(servers ...string) []string {
	return servers
}
