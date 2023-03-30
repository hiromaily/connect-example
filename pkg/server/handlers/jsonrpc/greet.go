package jsonrpc

//go:generate zenrpc

import (
	"context"
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/usecases/greet"
	"github.com/semrush/zenrpc/v2"
)

type GreetService struct {
	zenrpc.Service
	logger  logger.Logger
	ucGreet *greet.Greet
}

func NewGreetService(logger logger.Logger, ucGreet *greet.Greet) GreetService {
	return GreetService{
		logger:  logger,
		ucGreet: ucGreet,
	}
}

// Greet
func (g GreetService) Greet(ctx context.Context, name string) (string, *zenrpc.Error) {
	//r, _ := zenrpc.RequestFromContext(ctx)

	reply := g.ucGreet.Greet(name)

	return reply, nil
}
