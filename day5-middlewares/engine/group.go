package engine

import "log"

type Group struct {
	prefix      string
	parent      *Group
	engine      *Engine
	middlewares []HandlerFunc
}

func (g *Group) NewGroup(prefix string) *Group {
	e := g.engine
	newGroup := &Group{
		prefix: g.prefix + prefix,
		parent: g,
		engine: e,
	}
	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (g *Group) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	// g.engine.addRoute(method, pattern, handler) 内嵌导致的模糊
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *Group) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *Group) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}
