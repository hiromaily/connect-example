package greet

import (
	"fmt"

	"github.com/hiromaily/connect-example/pkg/logger"
)

type Greet struct {
	logger logger.Logger
}

func NewUseCaseGreet(logger logger.Logger) *Greet {
	return &Greet{
		logger: logger,
	}
}

// Greet returns reply
func (g *Greet) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
