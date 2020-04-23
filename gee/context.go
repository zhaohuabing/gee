package gee
import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Context struct{
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
}

func newContext(Writer http.ResponseWriter, Req *http.Request) *Context{
	return &Context{
		Writer: Writer,
		Req: Req,
		Path: Req.URL.Path,
		Method: Req.Method,
		StatusCode:400,
	}
}

func (c *Context) PostForm(key string) string{
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int){
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string){
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ... interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}){
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	fmt.Printf("%Tssssss",obj)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj);err !=nil{
		http.Error(c.Writer, err.Error(),500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}