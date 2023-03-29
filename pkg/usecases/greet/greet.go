package greet

import (
	"fmt"

	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/repositories"
)

type Greet struct {
	logger    logger.Logger
	greetRepo repositories.GreetTableRepo
}

func NewUseCaseGreet(logger logger.Logger, greetRepo repositories.GreetTableRepo) *Greet {
	return &Greet{
		logger:    logger,
		greetRepo: greetRepo,
	}
}

// Greet returns reply
func (g *Greet) Greet(name string) string {
	greet := g.greetRepo.GetGreet()
	return fmt.Sprintf("%s, %s!", greet, name)
}
