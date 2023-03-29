package eliza

import (
	"fmt"

	"github.com/hiromaily/connect-example/pkg/logger"
)

type Eliza struct {
	logger logger.Logger
}

func NewUseCaseEliza(logger logger.Logger) *Eliza {
	return &Eliza{
		logger: logger,
	}
}

// Say returns words
func (e *Eliza) Say(sentence string) string {
	return fmt.Sprintf("You said, %s!", sentence)
}

func (e *Eliza) GetIntros(name string) []string {
	return []string{"Hello", "Goodbye", "How are you"}
}
