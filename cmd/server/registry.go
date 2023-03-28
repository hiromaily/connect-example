package main

import (
	"github.com/hiromaily/connect-example/pkg/server"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/hiromaily/connect-example/pkg/apis"
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/server/cors"
)

type Registry interface {
	NewServer() server.Server
}

type registory struct {
	mux    *http.ServeMux
	logger logger.Logger
}

func NewRegistory() Registry {
	reg := &registory{}
	reg.mux = http.NewServeMux()
	reg.newLogger()

	return reg
}

func (r *registory) NewServer() server.Server {
	// switch server by DI
	return r.newConnectServer()
}

func (r *registory) newLogger() logger.Logger {
	if r.logger == nil {
		r.logger = logger.NewZeroLog()
	}
	return r.logger
}

func (r *registory) newConnectServer() server.Server {
	r.createHandlers()

	addr := "localhost:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	srv := &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			cors.NewCORS().Handler(r.mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	return server.NewServer(srv, r.newLogger())
}

func (r *registory) createHandlers() {
	// params: path, handler
	r.mux.Handle(apis.NewGreetHandler(r.newLogger()))
	r.mux.Handle(apis.NewElizaHandler(r.newLogger()))
}
