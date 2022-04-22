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

func testHandler_TemlateData_Map(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.html"))

	t.ExecuteTemplate(w, "main.html", map[string]interface{}{
		"Title":  "Template",
		"Header": "Hello World!",
	})

}

type Address struct {
	City   string
	Street string
}

type Page struct {
	Title  string
	Header string
	Adress Address
}

func testHandler_TemlateData_Struct(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.html"))

	t.ExecuteTemplate(w, "main.html", Page{
		Title:  "Template Struct",
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

	testHandler_TemlateData_Struct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
