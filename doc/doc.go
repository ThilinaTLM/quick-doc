package doc

// UiTheme UI types
type UiTheme int

const (
	UI_THEME_RAPI_DOC   = UiTheme(iota)
	UI_THEME_SWAGGER_UI = UiTheme(iota)
)

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

// Config API Doc configuration
type Config struct {
	Enabled     bool
	Title       string
	Description string
	Version     string
	Servers     []string
	AuthTypes   []AuthType
	SpecUrl     string
	UiEnabled   bool
	UiUrl       string
	UiTheme     UiTheme
}

type Endpoint struct {
	method      MethodType
	Path        string
	Summary     string
	Description string
	AuthTypes   []AuthType
	RequestBody RequestBody
	Tags        []string
	QueryParams Parameters
	PathParams  Parameters
	Headers     Parameters
	Responses   Responses
}

type Doc struct {
	config    Config
	endpoints []Endpoint
}

func NewDoc(config Config) *Doc {
	return &Doc{
		config:    config,
		endpoints: make([]Endpoint, 0),
	}
}

func (d *Doc) addEndpoint(ep Endpoint) {
	if ep.AuthTypes == nil {
		ep.AuthTypes = make([]AuthType, 0)
	}
	d.endpoints = append(d.endpoints, ep)
}

// Get add get endpoint
func (d *Doc) Get(ep Endpoint) {
	ep.method = METHOD_GET
	d.addEndpoint(ep)
}

// Post add post endpoint
func (d *Doc) Post(ep Endpoint) {
	ep.method = METHOD_POST
	d.addEndpoint(ep)
}

// Put add put endpoint
func (d *Doc) Put(ep Endpoint) {
	ep.method = METHOD_PUT
	d.addEndpoint(ep)
}

// Delete add delete endpoint
func (d *Doc) Delete(ep Endpoint) {
	ep.method = METHOD_DELETE
	d.addEndpoint(ep)
}

func Tags(tags ...string) []string {
	return tags
}
