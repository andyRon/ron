package ronweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context （上下文），封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
}

// Context的构造器
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// Status 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key string, val string) {
	c.W.Header().Set(key, val)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.W.Write([]byte(html))
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W) // TODO
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
