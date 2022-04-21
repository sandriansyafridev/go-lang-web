package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequest(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home page")
		fmt.Fprintln(w, "METHOD: ", r.Method) // [GET,POST,PUT,PATCH,DELETE]
		fmt.Fprintln(w, "URL: ", r.URL)       // url

	})

	http.ListenAndServe(":1234", mux)

}
