package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hiromaily/connect-example/pkg/logger"
)

// Refer to
// https://github.com/bufbuild/connect-demo/blob/3a30d4de07d6ac42110acd4ebf64bb4bf8a62579/main.go

func main() {
	//TODO: config file
	reg := NewRegistory()
	srv := reg.NewServer()
	logger := reg.NewLogger()
	startServer(srv, logger)
}

func startServer(srv *http.Server, logger logger.Logger) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	logger.Info(fmt.Sprintf("start server: %s", srv.Addr))
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Errorf("fail to linten and serve: %v", err))
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Errorf("fail to shutdown server: %v", err)) //nolint:gocritic
	}
}
