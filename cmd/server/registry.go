package main

import (
	"github.com/hiromaily/connect-example/pkg/server/handlers/jsonrpc"
	"net/http"
	"os"
	"time"

	"github.com/semrush/zenrpc/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/repositories"
	"github.com/hiromaily/connect-example/pkg/server"
	"github.com/hiromaily/connect-example/pkg/server/cors"
	connect2 "github.com/hiromaily/connect-example/pkg/server/handlers/connect"
	"github.com/hiromaily/connect-example/pkg/storages/mock"
	"github.com/hiromaily/connect-example/pkg/usecases/eliza"
	"github.com/hiromaily/connect-example/pkg/usecases/greet"
)

type Registry interface {
	NewConnectServer() server.Server
	NewJSONRPCServer() server.Server
}

type registory struct {
	logger    logger.Logger
	ucGreet   *greet.Greet
	ucEliza   *eliza.Eliza
	greetRepo repositories.GreetTableRepo
	elizaRepo repositories.ElizaTableRepo
}

func NewRegistory() Registry {
	reg := &registory{}
	reg.newLogger()

	return reg
}

func (r *registory) NewConnectServer() server.Server {
	// switch server by DI or call another NewXXXServer()
	return r.newConnectServer()
}

func (r *registory) NewJSONRPCServer() server.Server {
	// switch server by DI or call another NewXXXServer()
	return r.newJSONRPCServer()
}

func (r *registory) newLogger() logger.Logger {
	if r.logger == nil {
		// switch logger
		r.logger = logger.NewZeroLog()
	}
	return r.logger
}

func (r *registory) newUseCaseGreet() *greet.Greet {
	if r.ucGreet == nil {
		r.ucGreet = greet.NewUseCaseGreet(r.newLogger(), r.newGreetRepo())
	}
	return r.ucGreet
}

func (r *registory) newUseCaseEliza() *eliza.Eliza {
	if r.ucEliza == nil {
		r.ucEliza = eliza.NewUseCaseEliza(r.newLogger(), r.newElizaRepo())
	}
	return r.ucEliza
}

func (r *registory) newGreetRepo() repositories.GreetTableRepo {
	if r.greetRepo == nil {
		// switch Database at second parameter
		r.greetRepo = repositories.NewGreetTableRepo(r.newLogger(), mock.NewGreetTable(r.newLogger()))
	}
	return r.greetRepo
}

func (r *registory) newElizaRepo() repositories.ElizaTableRepo {
	if r.elizaRepo == nil {
		// switch Database at second parameter
		r.elizaRepo = repositories.NewElizaTableRepo(r.newLogger(), mock.NewElizaTable(r.newLogger()))
	}
	return r.elizaRepo
}

func (r *registory) newConnectServer() server.Server {
	mux := http.NewServeMux()
	r.createConnectHandlers(mux)

	addr := "localhost:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	srv := &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			cors.NewCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	return server.NewServer(srv, r.newLogger())
}

func (r *registory) newJSONRPCServer() server.Server {
	mux := http.NewServeMux()
	r.createJSONRPCHandlers(mux)

	addr := "localhost:8090"
	if port := os.Getenv("JSONRPC_PORT"); port != "" {
		addr = ":" + port
	}
	// must be HTTP/1.1
	srv := &http.Server{
		Addr:              addr,
		Handler:           cors.NewCORS().Handler(mux),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	return server.NewServer(srv, r.newLogger())
}

func (r *registory) createConnectHandlers(mux *http.ServeMux) {
	// params: path, handler
	mux.Handle(connect2.NewGreetHandler(r.newLogger(), r.newUseCaseGreet()))
	mux.Handle(connect2.NewElizaHandler(r.newLogger(), r.newUseCaseEliza()))
	mux.Handle(connect2.NewJSONRPCHandler(r.newLogger(), r.newUseCaseGreet()))
}

func (r *registory) createJSONRPCHandlers(mux *http.ServeMux) {
	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	//rpc.Register("arith", testdata.ArithService{})
	//rpc.Register("", testdata.ArithService{}) // public
	rpc.Register("arith", jsonrpc.ArithService{})
	rpc.Register("greet", jsonrpc.NewGreetService(r.newLogger(), r.newUseCaseGreet()))

	//rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	mux.Handle("/", rpc)
}
