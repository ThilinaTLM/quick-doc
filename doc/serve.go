package doc

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (cd *CompiledDoc) ServeHttp(r *mux.Router) {
	cd.ServeHttpJson(r)
	cd.ServeHttpUi(r)
}

func (cd *CompiledDoc) ServeHttpJson(r *mux.Router) {
	r.HandleFunc(cd.raw.config.SpecUrl, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(cd.Json)
		if err != nil {
			return
		}
	})
}

func (cd *CompiledDoc) ServeHttpUi(r *mux.Router) {
	if !cd.raw.config.UiEnabled {
		return
	}
	var html string
	if cd.raw.config.UiTheme == UI_THEME_RAPI_DOC {
		html = rapiDocDark(cd.raw.config.Title, cd.raw.config.SpecUrl)
	} else if cd.raw.config.UiTheme == UI_THEME_SWAGGER_UI {
		html = swaggerUiDefault(cd.raw.config.Title, cd.raw.config.SpecUrl)
	} else {
		panic(fmt.Sprintf("Unknown UI theme: %v", cd.raw.config.UiTheme))
		return
	}

	r.HandleFunc(cd.raw.config.UiUrl, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(html))
		if err != nil {
			return
		}
	})
}

func rapiDocDark(title string, specUrl string) string {
	return fmt.Sprintf(`
			<!doctype html>
			<html>
			  <head>
				<meta charset="utf-8">
				<title>%s</title>
				<script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script>
			  </head>
			  <body>
				<rapi-doc
				  spec-url = "%s"
				> </rapi-apiDoc>
			  </body>
			</html>
		`, title, specUrl)
}

func swaggerUiDefault(title string, specUrl string) string {
	return fmt.Sprintf(`
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<meta http-equiv="X-UA-Compatible" content="ie=edge">
					<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-standalone-preset.js"></script>
					<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-bundle.js"></script>
					<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui.css" />
					<title>%s</title>
					<style>
						body {
							margin: 0;
							padding: 0;
						}
					</style>
				</head>
				<body>
					<div id="swagger-ui"></div>
					<script>
						window.onload = function() {
						  SwaggerUIBundle({
							url: "%s",
							dom_id: '#swagger-ui',
							presets: [
							  SwaggerUIBundle.presets.apis,
							  SwaggerUIStandalonePreset
							],
							layout: "StandaloneLayout"
						  })
						}
					</script>
				</body>
				</html>
			`, title, specUrl)
}
