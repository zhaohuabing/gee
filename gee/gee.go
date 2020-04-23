package gee

import (
	"net/http"
)

type Engine struct{
	router * Router
}

func New() *Engine{
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func(engine *Engine) GET(pattern string, handler HandlerFunc){
	engine.addRoute("GET",pattern, handler)
}
func(engine *Engine) POST(pattern string, handler HandlerFunc){
	engine.addRoute("POST",pattern, handler)
}
func(engine *Engine) PUT(pattern string, handler HandlerFunc){
	engine.addRoute("PUT",pattern, handler)
}

func(engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	engine.router.handle(newContext(writer,req))
}

func(engine *Engine) Run(addr string) error{
	return http.ListenAndServe(addr, engine)
}
