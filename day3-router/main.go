package main

import (
	"gee/engine"
	"net/http"
)

func main() {
	e := engine.New()
	e.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	e.GET("/hello", func(c *engine.Context) {
		// expect /hello?name=Qian
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	e.GET("/hello/:name", func(c *engine.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Params["name"], c.Path)
	})
	e.GET("/assets/*filepath", func(c *engine.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"filepath": c.Params["filepath"],
		})
	})
	e.Run(":8080")
}
