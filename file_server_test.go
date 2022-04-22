package golangweb

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func testHandler_SimpleHTML_Directory(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseGlob("./template/*.gohtml"))
	t.ExecuteTemplate(w, "main.gohtml", "Hello")

}

func TestSimpleHTML_Directory(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_SimpleHTML_Directory(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
