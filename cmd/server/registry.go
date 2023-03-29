package main

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/hiromaily/connect-example/pkg/apis/connect"
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/server"
	"github.com/hiromaily/connect-example/pkg/server/cors"
	"github.com/hiromaily/connect-example/pkg/usecases/eliza"
	"github.com/hiromaily/connect-example/pkg/usecases/greet"
)

type Registry interface {
	NewServer() server.Server
}

type registory struct {
	mux     *http.ServeMux
	logger  logger.Logger
	ucGreet *greet.Greet
	ucEliza *eliza.Eliza
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

func (r *registory) newUseCaseGreet() *greet.Greet {
	if r.ucGreet == nil {
		r.ucGreet = greet.NewUseCaseGreet(r.newLogger())
	}
	return r.ucGreet
}

func (r *registory) newUseCaseEliza() *eliza.Eliza {
	if r.ucEliza == nil {
		r.ucEliza = eliza.NewUseCaseEliza(r.newLogger())
	}
	return r.ucEliza
}

func (r *registory) newConnectServer() server.Server {
	r.createConnectHandlers()

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

	return server.NewConnectServer(srv, r.newLogger())
}

func (r *registory) createConnectHandlers() {
	// params: path, handler
	r.mux.Handle(connect.NewGreetHandler(r.newLogger(), r.newUseCaseGreet()))
	r.mux.Handle(connect.NewElizaHandler(r.newLogger(), r.newUseCaseEliza()))
}
