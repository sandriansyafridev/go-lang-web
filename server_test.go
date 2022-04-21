package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerHTTP(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home page")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About page")
	})

	http.ListenAndServe(":1234", mux)

}
