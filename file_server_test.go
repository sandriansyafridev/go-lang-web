package golangweb

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func testHandler_SimpleHTML_File(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./template/main.html"))

	t.ExecuteTemplate(w, "main.html", "Hello World!")

}

func TestSimpleHTML_File(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:1234/", nil)
	recorder := httptest.NewRecorder()

	testHandler_SimpleHTML_File(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
