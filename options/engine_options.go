package options

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine  *gin.Engine
	host    string
	port    string
	timeout time.Duration
}

func ServerOptions(options ...func(*Engine)) *Engine {
	engine := &Engine{
		engine: gin.Default(),
	}
	for _, e := range options {
		e(engine)
	}
	return engine
}

func (e *Engine) Start() error {
	return e.engine.Run(e.host + ":" + e.port)
}

func WithHost(host string) func(*Engine) {
	return func(eo *Engine) {
		eo.host = host
	}
}

func WithPort(port string) func(*Engine) {
	return func(eo *Engine) {
		eo.port = port
	}
}

func WithRoutes(routeFuncs ...func(*gin.RouterGroup)) func(*Engine) {
	return func(e *Engine) {
		v1 := e.engine.Group("v1")
		for _, route := range routeFuncs {
			route(v1)
		}
	}
}


func WithCorsRules(corsConfig cors.Config) func(*Engine) {
	return func(e *Engine) {
		corsMiddleware := cors.New(corsConfig)
		e.engine.Use(corsMiddleware)
	}
}
