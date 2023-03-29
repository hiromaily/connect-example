package eliza

import (
	"fmt"

	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/repositories"
)

type Eliza struct {
	logger    logger.Logger
	elizaRepo repositories.ElizaTableRepo
}

func NewUseCaseEliza(logger logger.Logger, elizaRepo repositories.ElizaTableRepo) *Eliza {
	return &Eliza{
		logger:    logger,
		elizaRepo: elizaRepo,
	}
}

// Say returns words
func (e *Eliza) Say(sentence string) string {
	word := e.elizaRepo.GetEliza()
	return fmt.Sprintf("%s, %s!", word, sentence)
}

func (e *Eliza) GetIntros(name string) []string {
	return []string{"Hello", "Goodbye", "How are you"}
}
