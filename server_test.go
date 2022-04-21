package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var target = "http://localhost:1234"

func testHandlerCreateHeader(w http.ResponseWriter, r *http.Request) {

	contentType_Header := r.Header.Get("content-type")
	fmt.Fprint(w, "Content-Type: ", contentType_Header)

}

func testHandlerGetHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("x-user", "Sandrian Syafri")
	fmt.Fprint(w, "TEST OK")

}

func TestCreateHeaderToServer(t *testing.T) {

	//create http test
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()

	//add header
	req.Header.Add("content-type", "application/json")

	//process testing
	testHandlerCreateHeader(rec, req)

	//get response after test
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	//log to console
	log.Println(string(body))

}

func TestGetHeaderFromClient(t *testing.T) {

	//create http test
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()

	//get header

	//process testing
	testHandlerGetHeader(rec, req)

	//get response after test
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	xUser_Header := res.Header.Get("x-user")

	//log to console
	log.Println(string(body))
	log.Println("X-USER: ", xUser_Header)

}
