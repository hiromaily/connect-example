package repositories

import (
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/storages"
)

// repositories are Interface Adapter for Database gateway

type ElizaTableRepo interface {
	GetEliza() string
	PutEliza(name string) error
}

type elizaTableRepo struct {
	logger        logger.Logger
	elizaStorager storages.ElizaStorager
}

func NewElizaTableRepo(logger logger.Logger, elizaStorager storages.ElizaStorager) ElizaTableRepo {
	return &elizaTableRepo{
		logger:        logger,
		elizaStorager: elizaStorager,
	}
}

func (e *elizaTableRepo) GetEliza() string {
	return e.elizaStorager.GetEliza()
}

func (e *elizaTableRepo) PutEliza(name string) error {
	return e.elizaStorager.PutEliza(name)
}
