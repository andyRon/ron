package main

import (
	"net/http"
	"ronweb/ronweb"
)

func main() {

	r := ronweb.New()
	r.GET("/", func(c *ronweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ron-web</h1>")
	})

	r.GET("/hello", func(c *ronweb.Context) {
		// /hello?name=andyron
		c.String(http.StatusOK, "hello %s, your paht is:%s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ronweb.Context) {
		c.JSON(http.StatusOK, ronweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/assets/*filepath", func(c *ronweb.Context) {
		c.JSON(http.StatusOK, ronweb.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
