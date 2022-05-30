package main

import (
	"fmt"
	"github.com/ThilinaTLM/quick-doc/qdoc"
	"github.com/ThilinaTLM/quick-doc/ui"
	"net/http"
)

type Project struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Description2 string `json:"description2"`
}

type ReqUserAdd struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Project *Project `json:"project"`
}

func main() {
	Doc()
}

func Doc() {
	doc := qdoc.NewDoc(qdoc.Config{
		Title:       "Quick Doc Demo",
		Description: "Quick Doc demo API documentation example",
		Version:     "1.0.0",
		Servers: qdoc.Servers(
			"http://localhost:8080",
			"http://dev.quickdoc.com",
		),
		SpecPath: "/doc/json",

		UiConfig: qdoc.UiConfig{
			Enabled:      true,
			Path:         "/doc/ui",
			DefaultTheme: ui.SWAGGER_UI,
			ThemeByQuery: false,
		},
	})

	doc.Post(&qdoc.Endpoint{
		Path: "/doc/user",
		Desc: "Create a new user",
		ReqBody: qdoc.ReqJson(doc.Schema(ReqUserAdd{
			Name: "Student 1",
			Age:  16,
			Project: &Project{
				Name:        "Volunteer Project",
				Description: "This is a volunteer project",
			},
		})),
		RespSet: qdoc.RespSet{
			Success: qdoc.ResJson("User creation success", nil),
		},
	}).Tag("User").WithBearerAuth()

	cd, err := doc.Compile()
	if err != nil {
		panic(err)
	}

	s := cd.ServeMux()

	fmt.Println("Server is running on port 8080")
	fmt.Println("Swagger UI: http://localhost:8080/doc/ui")
	err = http.ListenAndServe(":8080", s)
	if err != nil {
		return
	}
}
