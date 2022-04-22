package golangweb

import (
	"embed"
	"net/http"
	"testing"
)

//go:embed assets
var assets embed.FS

func testHandlerServeFile(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	switch {
	case name != "":
		http.ServeFile(w, r, "assets/index.html")
	default:
		http.ServeFile(w, r, "assets/not-found.html")
	}

}

func TestServeFile(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandlerServeFile)

	http.ListenAndServe(":1234", mux)

}
