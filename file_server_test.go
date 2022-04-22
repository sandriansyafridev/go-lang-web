package golangweb

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func testHandler_TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.html",
		"./templates/main.html",
		"./templates/content.html",
		"./templates/footer.html",
	))

	t.ExecuteTemplate(w, "main.html", map[string]interface{}{
		"Title":  "Template layout",
		"Header": "Hello World!",
		"Items": []string{
			"Item 1",
			"Item 2",
			"Item 3",
			"Item 4",
		},
	})
}

func TestTemplateLayout(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
