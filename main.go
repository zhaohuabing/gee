package main

import (
	"gee"
	"log"
	"net/http"
)
func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	engine.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, struct{
			username  string
			password string
		}{
			username : c.PostForm("username"),
			password: c.PostForm("password"),
		})
	})

	log.Fatal(engine.Run(":9990"));
}
