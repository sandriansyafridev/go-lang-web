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

	t = template.Must(template.ParseFS(templates, "templates/*.html"))
)

func testHandler_Autoescape(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"Title":  "Autoescape",
		"Header": "<p>Test Ting</p>",
	}
	t.ExecuteTemplate(w, "index", data)

}

func TestAutoescape(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_Autoescape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}

func TestAutoescapeOnBrowser(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler_Autoescape)

	http.ListenAndServe(":1234", mux)

}
