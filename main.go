package main

import (
	"net/http"
	"ron/ron"
)

func main() {

	r := ron.New()
	r.GET("/", func(c *ron.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ron</h1>")
	})

	r.GET("/hello", func(c *ron.Context) {
		// /hello?name=andyron
		c.String(http.StatusOK, "hello %s, your paht is:%s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ron.Context) {
		c.JSON(http.StatusOK, ron.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
