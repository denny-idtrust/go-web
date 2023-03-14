package main

import (
	"fmt"
	"net/http"
)

func main() {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			return
		}
		_, errs := fmt.Fprintf(w, "Hello World")
		if errs != nil {
			return
		}
	}

	server := http.Server{
		Addr:    "localhost:8181",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
