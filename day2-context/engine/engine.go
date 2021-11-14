package engine

import "net/http"

type HandlerFunc func(c *Context)

// 直接内嵌*router, 在main包还是用不了这些方法
type Engine struct {
	*router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.handle(c)
}
