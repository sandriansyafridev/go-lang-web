package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed assets
var assets embed.FS

func TestFileServer(t *testing.T) {

	dir, _ := fs.Sub(assets, "assets")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe(":1234", mux)
}
