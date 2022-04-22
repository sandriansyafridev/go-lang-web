package golangweb

import (
	"embed"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.html
var templates embed.FS

type Address struct {
	City   string
	Street string
}

type Page struct {
	Title  string
	Header string
	Adress Address
}

func testHandler_TemlateData_IfElse(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.html"))

	t.ExecuteTemplate(w, "main.html", Page{
		Title:  "Template IfElse",
		Header: "Hello World!",
		Adress: Address{
			Street: "Jl.Guru",
			City:   "Jambi",
		},
	})

}

func TestTemplateData_Struct(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_TemlateData_IfElse(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
