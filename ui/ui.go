package ui

import "fmt"

type Theme string

const (
	RAPI_DOC   Theme = "rapi-doc"
	SWAGGER_UI Theme = "swagger-ui"
)

type Config struct {
	Theme   Theme
	Title   string
	SpecUrl string
}

func (c *Config) Validate() error {
	if c.Theme != RAPI_DOC && c.Theme != SWAGGER_UI {
		return fmt.Errorf("invalid type: %s", c.Theme)
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

	if config.Theme == RAPI_DOC {
		return RapiDocHTML(config), nil
	} else if config.Theme == SWAGGER_UI {
		return SwaggerUiHTML(config), nil
	}
	return "", fmt.Errorf("invalid type: %s", config.Theme)
}
