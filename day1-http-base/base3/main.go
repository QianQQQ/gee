package main

import (
	"fmt"
	"net/http"
)

func main() {
	e := New()
	e.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	e.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	e.Run(":8080")
}

type Engine struct {
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{router: map[string]http.HandlerFunc{}}
}

func (e *Engine) addRoute(method, pattern string, handler http.HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler http.HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler http.HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
