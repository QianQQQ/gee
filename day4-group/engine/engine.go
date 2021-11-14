package engine

import "net/http"

type HandlerFunc func(c *Context)

// 直接内嵌*router, 在main包还是用不了这些方法
type Engine struct {
	*router
	*RouterGroup
	groups []*RouterGroup
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.handle(c)
}
