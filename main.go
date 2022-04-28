package main

import (
	"fmt"
	"github.com/ThilinaTLM/quick-doc/doc"
	"github.com/gorilla/mux"
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
	apiDoc := doc.NewDoc(doc.Config{
		Enabled:     true,
		Title:       "Quick Doc Demo",
		Description: "Quick doc demo API documentation example",
		Version:     "1.0.0",
		Servers:     []string{"http://localhost:8080"},
		AuthTypes:   doc.AuthTypesBearer(),
		SpecUrl:     "/api/doc/json",
		UiEnabled:   true,
		UiUrl:       "/api/doc/ui",
		UiTheme:     doc.UI_THEME_SWAGGER_UI,
	})

	apiDoc.Post(doc.Endpoint{
		Path:        "/api/user",
		Description: "Create a new user",
		Tags:        doc.Tags("User"),
		RequestBody: doc.ReqBodyJson(&ReqUserAdd{
			Project: nil,
		}),
		Responses: doc.Resp(
			doc.RespSuccess("User created successfully", nil),
		),
	})

	compiledDoc, err := apiDoc.Compile()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	compiledDoc.ServeHttp(router)

	fmt.Println("Server is running on port 8080")
	fmt.Println("Swagger UI: http://localhost:8080/api/doc/ui")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
