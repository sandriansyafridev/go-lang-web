package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandlerSingleQueryParams(w http.ResponseWriter, r *http.Request) {
	queryParams_Name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Query Params Name: ", queryParams_Name)
}

func testHandlerMultipleQueryParams(w http.ResponseWriter, r *http.Request) {
	queryParams_Name := r.URL.Query().Get("name")
	queryParams_Age := r.URL.Query().Get("age")
	fmt.Fprintln(w, "Query Params Name: ", queryParams_Name)
	fmt.Fprintln(w, "Query Params Age: ", queryParams_Age)
}

var targetSingleQueryParams = "http://localhost:1234/test-handler?name=Sandrian"
var targetMultipleQueryParams = "http://localhost:1234/test-handler?name=Sandrian&age=18"

func TestQueryParams(t *testing.T) {

	//create http test
	request := httptest.NewRequest("GET", targetSingleQueryParams, nil)
	recorder := httptest.NewRecorder()

	//run test
	testHandlerSingleQueryParams(recorder, request)

	//get response test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	//log  response test to console
	log.Println(string(body))

}

func TestMultipleQueryParams(t *testing.T) {

	//create http test
	request := httptest.NewRequest("GET", targetMultipleQueryParams, nil)
	recorder := httptest.NewRecorder()

	//run test
	testHandlerMultipleQueryParams(recorder, request)

	//get response test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	//log  response test to console
	log.Println(string(body))

}
