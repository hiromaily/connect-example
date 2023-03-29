package mock

import (
	"math/rand"

	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/storages"
)

type elizaTable struct {
	logger    logger.Logger
	elizaList []string
}

func NewElizaTable(logger logger.Logger) storages.ElizaStorager {
	return &elizaTable{
		logger:    logger,
		elizaList: []string{"You said", "Speaking"},
	}
}

func (e *elizaTable) GetEliza() string {
	idx := rand.Intn(len(e.elizaList))
	return e.elizaList[idx]
}

func (e *elizaTable) PutEliza(name string) error {
	e.elizaList = append(e.elizaList, name)
	return nil
}
