package qdoc

import (
	"fmt"
	"github.com/ThilinaTLM/quick-doc/ui"
	"net/http"
)

func serveJson(json []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(json)
		if err != nil {
			return
		}
	}
}

func serveUi(html string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(html))
		if err != nil {
			return
		}
	}
}

func serveUiDynamic(defaultTheme ui.Theme, htmlMap map[ui.Theme]string) func(http.ResponseWriter, *http.Request) {

	if _, ok := htmlMap[defaultTheme]; !ok {
		panic(fmt.Sprintf("Default theme %s not found", defaultTheme))
	}

	defaultHTML := htmlMap[defaultTheme]

	return func(w http.ResponseWriter, r *http.Request) {
		theme := r.URL.Query().Get("theme")
		if theme == "" {
			theme = string(defaultTheme)
		}
		var html string
		if v, ok := htmlMap[ui.Theme(theme)]; ok {
			html = v
		} else {
			html = defaultHTML
		}

		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(html))
		if err != nil {
			return
		}
	}
}

func (cd *CompiledDoc) ServeMux() *http.ServeMux {
	s := http.NewServeMux()
	s.HandleFunc(cd.config.SpecPath, serveJson(cd.Json))

	if cd.config.UiConfig.Enabled {
		if cd.config.UiConfig.ThemeByQuery {
			var htmlMap = make(map[ui.Theme]string)
			for _, theme := range []ui.Theme{ui.SWAGGER_UI, ui.RAPI_DOC} {
				html, err := ui.HTML(ui.Config{
					Theme:   theme,
					Title:   cd.config.Title,
					SpecUrl: cd.config.SpecPath,
				})
				if err != nil {
					panic(err)
				}
				htmlMap[theme] = html
			}

			s.HandleFunc(cd.config.UiConfig.Path, serveUiDynamic(cd.config.UiConfig.DefaultTheme, htmlMap))
		} else {
			html, err := ui.HTML(ui.Config{
				Theme:   cd.config.UiConfig.DefaultTheme,
				Title:   cd.config.Title,
				SpecUrl: cd.config.SpecPath,
			})
			if err != nil {
				panic(err)
			}
			s.HandleFunc(cd.config.UiConfig.Path, serveUi(html))
		}

	}

	return s
}
