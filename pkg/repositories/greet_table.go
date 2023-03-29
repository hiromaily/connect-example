package repositories

import (
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/storages"
)

// repositories are Interface Adapter for Database gateway

type GreetTableRepo interface {
	GetGreet() string
	PutGreet(name string) error
}

type greetTableRepo struct {
	logger        logger.Logger
	greetStorager storages.GreetStorager
}

func NewGreetTableRepo(logger logger.Logger, greetStorager storages.GreetStorager) GreetTableRepo {
	return &greetTableRepo{
		logger:        logger,
		greetStorager: greetStorager,
	}
}

func (g *greetTableRepo) GetGreet() string {
	return g.greetStorager.GetGreet()
}

func (g *greetTableRepo) PutGreet(name string) error {
	return g.greetStorager.PutGreet(name)
}
