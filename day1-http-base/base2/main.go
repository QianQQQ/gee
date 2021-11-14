package main

import (
	"fmt"
	"log"
	"net/http"
)

// 引擎用来处理request, 就是一个实现了ServeHTTP的handler
type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &Engine{}))
}
