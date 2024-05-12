package ron

import (
	"log"
	"net/http"
)

// 路由
type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, patten string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, patten)
	key := method + "-" + patten
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusOK, "404 NOT FOUND: $s\n", key)
	}

}
