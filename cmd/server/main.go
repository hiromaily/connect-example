package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/hiromaily/connect-example/pkg/apis"
	"github.com/hiromaily/connect-example/pkg/server/cors"
)

// Refer to
// https://github.com/bufbuild/connect-demo/blob/3a30d4de07d6ac42110acd4ebf64bb4bf8a62579/main.go

func main() {
	mux := http.NewServeMux()
	createHandlers(mux)

	startServer(mux)
}

func createHandlers(mux *http.ServeMux) {
	// params: path, handler
	mux.Handle(apis.NewGreetHandler())
	mux.Handle(apis.NewElizaHandler())
}

func startServer(mux *http.ServeMux) {
	addr := "localhost:8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	// serve
	//http.ListenAndServe(
	//	"localhost:8080",
	//	// h2c protocol is the non-TLS version of HTTP/2
	//	h2c.NewHandler(mux, &http2.Server{}),
	//)

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
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err) //nolint:gocritic
	}
}
