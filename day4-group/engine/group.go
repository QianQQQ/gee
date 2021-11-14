package engine

import "log"

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	engine      *Engine
	middlewares []HandlerFunc
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	e := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: e,
	}
	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	// g.engine.addRoute(method, pattern, handler) 内嵌导致的模糊
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}
