package golangweb

import (
	"embed"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	//go:embed templates/*.html
	templates embed.FS
	t         = template.Must(template.ParseFS(templates, "templates/*.html"))
)

func testHandler_TemplateCache(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":  "Template Cache",
		"Header": "Hello World!",
	}
	t.ExecuteTemplate(w, "main", data)
}

func TestTemplateFunction(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/test", nil)
	recorder := httptest.NewRecorder()

	testHandler_TemplateCache(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))
}
