package server

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

// duplicated: must be removed

type connectServer struct {
	srv    *http.Server
	logger logger.Logger
}

func NewConnectServer(srv *http.Server, logger logger.Logger) Server {
	return &connectServer{
		srv:    srv,
		logger: logger,
	}
}

func (c *connectServer) Run() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	c.logger.Info(fmt.Sprintf("start server: %s", c.srv.Addr))
	go func() {
		if err := c.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.logger.Fatal(fmt.Errorf("fail to linten and serve: %v", err))
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := c.srv.Shutdown(ctx); err != nil {
		c.logger.Fatal(fmt.Errorf("fail to shutdown server: %v", err)) //nolint:gocritic
	}
}
