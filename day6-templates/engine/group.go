package engine

import (
	"log"
	"net/http"
	"path"
)

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

// 映射到服务器上文件的真实地址
func (g *Group) Static(relativePath string, root string) {
	// 获取绝对地址
	absolutePath := path.Join(g.prefix, relativePath)
	// Dir是一个类型
	fileSystem := http.Dir(root)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fileSystem))
	urlPattern := path.Join(relativePath, "/*filepath")
	g.GET(urlPattern, func(c *Context) {
		file := c.Params["filepath"]
		if _, err := fileSystem.Open(file); err != nil {
			c.Status(404)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	})
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
