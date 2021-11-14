package engine

import (
	"gee/engine/trie"
	"strings"
)

type router struct {
	// roots["GET"] roots["Post"]
	roots map[string]*trie.Node
	// handlers["GET-/p/:lang/doc"]
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    map[string]*trie.Node{},
		handlers: map[string]HandlerFunc{},
	}
}

// 主要是截断第一个*之后的玩意, 还有把 a//b这种情况解决
func parsePattern(pattern string) (parts []string) {
	vs := strings.Split(pattern, "/")
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &trie.Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRouter(method, path string) (*trie.Node, map[string]string) {
	// 把输入路径切割
	searchParts := parsePattern(path)
	params := map[string]string{}
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	// 找到输入路径对应的Node
	n := root.Search(searchParts, 0)
	// 有对应的节点
	if n != nil {
		parts := parsePattern(n.Pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(404, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
