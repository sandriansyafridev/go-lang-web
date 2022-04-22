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

func testHandler_TemplateAction_Range(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.html"))

	t.ExecuteTemplate(w, "main.html", map[string]interface{}{
		"Items": []string{
			"Item 1", "Item 2", "Item 3", "Item 4",
		},
	})

}

func TestTemplateAction_Range(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_TemplateAction_Range(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
