package mock

import (
	"math/rand"

	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/storages"
)

type greetTable struct {
	logger    logger.Logger
	greetList []string
}

func NewGreetTable(logger logger.Logger) storages.GreetStorager {
	return &greetTable{
		logger:    logger,
		greetList: []string{"Hello", "Hey", "How's it going"},
	}
}

func (g *greetTable) GetGreet() string {
	idx := rand.Intn(len(g.greetList))
	return g.greetList[idx]
}

func (g *greetTable) PutGreet(name string) error {
	g.greetList = append(g.greetList, name)
	return nil
}
