package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var target = "http://localhost:1234"

func testHandlerPostForm(w http.ResponseWriter, r *http.Request) {

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	fmt.Fprintln(w, "Name:", name)
	fmt.Fprintln(w, "Email:", email)

}

func TestPostForm(t *testing.T) {

	bodyRequest := strings.NewReader("name=Sandrian Syafri&email=sandriansyafri@gmail.com")
	request := httptest.NewRequest(http.MethodPost, target+"/users", bodyRequest)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	testHandlerPostForm(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
