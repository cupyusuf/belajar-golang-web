package belajar_golang_web

import (
	"net/http"
	"testing"
)

func ServeFile(write http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(write, request, "./resources/ok.html")
		return
	} else {
		http.ServeFile(write, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
