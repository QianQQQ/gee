package main

import (
	"gee/engine"
	"gee/engine/middlewares"
	"log"
	"net/http"
	"time"
)

func onlyForV2() engine.HandlerFunc {
	return func(c *engine.Context) {
		t := time.Now()
		c.Fail(500, "test")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	e := engine.New()
	e.Use(middlewares.Logger())
	e.GET("/index", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := e.NewGroup("/v1")
	{
		v1.GET("/", func(c *engine.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *engine.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := e.NewGroup("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *engine.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *engine.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	e.Run(":8080")
}
