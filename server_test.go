package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerHTTP(t *testing.T) {

	var homeHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Home page.")) // manual
		fmt.Fprint(w, "Home")
	}

	var aboutHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About")
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":1234", nil)

}
