package gee

import "net/http"

type HandlerFunc func(ctx *Context)

type Router struct{
	handlers map[string]HandlerFunc
}

func newRouter() *Router{
	return & Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func(router *Router) addRoute(method string, pattern string,handler HandlerFunc){
key:=method +"_"+pattern
	router.handlers[key]=handler
}

func(router *Router)  handle(ctx *Context){
	key:= ctx.Method+"_"+ctx.Path
	if handler, ok := router.handlers[key]; ok {
		handler(ctx)
	}else{
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}