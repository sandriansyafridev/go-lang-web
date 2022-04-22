package golangweb

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	dir := http.Dir("./assets")

	fileServer := http.StripPrefix("/static/", http.FileServer(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/", fileServer)

	http.ListenAndServe(":1234", mux)
}
