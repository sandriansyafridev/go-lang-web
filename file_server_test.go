package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Person struct {
	Name string
}

func (p Person) SayHelloTo(name string) string {
	str := fmt.Sprintf("Hello %s, my name is %s", name, p.Name)
	return str
}

func testHandler_TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/main.html",
		"./templates/header.html",
		"./templates/content.html",
		"./templates/footer.html",
	))

	data := map[string]interface{}{
		"Title": "Function",
		"Person": Person{
			Name: "Sandrian",
		},
	}

	t.ExecuteTemplate(w, "main", data)
}

func TestTemplateFunction(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/test", nil)
	recorder := httptest.NewRecorder()

	testHandler_TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))
}
