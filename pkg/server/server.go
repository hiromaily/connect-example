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

type Server interface {
	Run()
}

type server struct {
	srv    *http.Server
	logger logger.Logger
}

func NewServer(srv *http.Server, logger logger.Logger) Server {
	return &server{
		srv:    srv,
		logger: logger,
	}
}

func (s *server) Run() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	s.logger.Info(fmt.Sprintf("start server: %s", s.srv.Addr))
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal(fmt.Errorf("fail to linten and serve: %v", err))
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Fatal(fmt.Errorf("fail to shutdown server: %v", err)) //nolint:gocritic
	}
}
