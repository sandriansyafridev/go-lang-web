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

func SetCookie(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{}
	cookie.Name = "x-com-user"
	cookie.Value = "Sandrian Syafri"
	cookie.Path = "/"

	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "Cookie Created!")
}
func GetCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("x-com-user")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Cookie %s: %s", cookie.Name, cookie.Value)
	}

}

func TestCookieOnBrowser(t *testing.T) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})
	http.HandleFunc("/get-cookie", GetCookie)
	http.HandleFunc("/set-cookie", SetCookie)

	http.ListenAndServe(":1234", nil)

}

func TestSetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, target+"/set-cookie", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}

func TestGetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, target+"/get-cookie", nil)
	cookie := http.Cookie{}
	cookie.Name = "x-com-user"
	cookie.Value = "Sandrian"

	request.AddCookie(&cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	log.Println(string(body))

}
