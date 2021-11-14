package engine

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

// 直接内嵌*router, 在main包还是用不了这些方法
type Engine struct {
	*router
	*Group
	groups []*Group
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.Group = &Group{engine: e}
	e.groups = []*Group{e.Group}
	return e
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	e.handle(c)
}
