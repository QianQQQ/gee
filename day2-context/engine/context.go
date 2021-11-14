package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 给用户提供更粗粒度的选项, 简化相关接口的调用
// 支撑额外的功能, 比如动态路由, 中间件
// 可以找到一次会话所需要的所有东西
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	// 请求的信息
	Method string
	Path   string
	// 响应的信息
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Method:  r.Method,
		Path:    r.URL.Path,
	}
}

// 拿到GET方法跟着的参数
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// 拿到POST方法的表单参数
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Header(k, v string) {
	c.Writer.Header().Set(k, v)
}

// text-plain
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Header("Content-Type", "text-plain")
	c.Status(code)
	// Writer的Write方法接收的是字节切片
	// 用Sprintf构造字符串, 然后利用字符串和字节切片的特性直接转换
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// application/json
func (c *Context) JSON(code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.Status(code)
	// NewEncoder returns a new encoder that writes to w.
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		// 原生的报错, String, JSON 方法参考了实现
		// SetHeader, StatusCode, body
		http.Error(c.Writer, err.Error(), 500)
	}
}

// text/html
func (c *Context) HTML(code int, html string) {
	c.Header("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
