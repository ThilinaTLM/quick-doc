package ui

import "fmt"

type Type string

const (
	RAPI_DOC   Type = "RAPI_DOC"
	SWAGGER_UI Type = "SWAGGER_UI"
)

type Config struct {
	Type    Type
	Theme   string
	Title   string
	SpecUrl string
}

func (c *Config) Validate() error {
	if c.Type != RAPI_DOC && c.Type != SWAGGER_UI {
		return fmt.Errorf("invalid type: %s", c.Type)
	}
	if c.Title == "" {
		return fmt.Errorf("title is required")
	}
	if c.SpecUrl == "" {
		return fmt.Errorf("spec url is required")
	}
	return nil
}

func HTML(config Config) (string, error) {
	err := config.Validate()
	if err != nil {
		return "", err
	}

	if config.Type == RAPI_DOC {
		return RapiDocHTML(config), nil
	} else if config.Type == SWAGGER_UI {
		return SwaggerUiHTML(config), nil
	}
	return "", fmt.Errorf("invalid type: %s", config.Type)
}
