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

//go:embed templates/*.gohtml
var templates embed.FS

func testHandler_SimpleHTML_Directory_Embed(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "main.gohtml", "Hello")

}

func TestSimpleHTML_DirectoryWithEmbed(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_SimpleHTML_Directory_Embed(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
