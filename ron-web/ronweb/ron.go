package ronweb

import (
	"net/http"
)

// HandlerFunc 用于定义路由映射的处理方法
type HandlerFunc func(c *Context)

// Engine 实现了ServeHTTP接口
type Engine struct {
	// 路由映射表，key由请求方法和静态路由地址构成，value是用户映射的处理方法
	router *router
}

// New 是Engine的构造器
func New() *Engine {
	return &Engine{router: NewRouter()}
}

// @param method string http访问方法
// @param pattern string http访问地址
// @param handler HandlerFunc 处理方法
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run 运行一个http服务器
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 解析请求的路径，查找路由映射表
// 如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(newContext(w, r))
}
