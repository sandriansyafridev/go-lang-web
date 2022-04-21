package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP Testing")
}

var target = "http://localhost:1234/test-handler"

func TestHTTPTest(t *testing.T) {

	//create http test
	request := httptest.NewRequest("GET", target, nil)
	recorder := httptest.NewRecorder()

	//run test
	testHandler(recorder, request)

	//get response test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	//log  response test to console
	log.Println(string(body))

}
