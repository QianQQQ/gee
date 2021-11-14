package engine

import (
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{map[string]HandlerFunc{}}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %s-%s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 接受的是一个context哦
func (r *router) handle(c *Context) {
	key := c.Request.Method + "-" + c.Request.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}
