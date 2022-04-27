package main

import (
	"fmt"
	"git.mytaxi.lk/pickme/delivery-services/quick-doc/doc"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ReqUserAdd struct {
	//Name    string      `json:"name"`
	//Age     int         `json:"age"`
	Project interface{} `json:"project"`
}

func main() {
	Test()
}

func Test() {
	obj := ReqUserAdd{}
	// print obj fields
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, v.Field(i).Type())
	}
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
		Tags:        []string{"user"},
		RequestBody: doc.ReqBodyJson(ReqUserAdd{}),
		Responses:   doc.Resp(doc.RespSuccess("User created successfully", nil)),
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
