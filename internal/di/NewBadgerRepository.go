package di

import (
	"github.com/sandronister/blockchain-go/configs"
	"github.com/sandronister/blockchain-go/internal/infra/database/connection/badger"
	badgerepository "github.com/sandronister/blockchain-go/internal/infra/database/repository/badger-repository"
	"github.com/sandronister/blockchain-go/internal/repository"
)

func NewBadgerRepository(cfg *configs.Config) (repository.IRepository, error) {
	connection, err := badger.GetConnection(cfg)
	if err != nil {
		return nil, err
	}
	return badgerepository.NewBadgerRepository(connection), nil
}
