package golangweb

import (
	"net/http"
	"testing"
)

func TestServerHTTP(t *testing.T) {

	server := http.Server{
		Addr: "localhost:1234",
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Error("Error: ", err.Error())
	}

}
