package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/entity"
	"github.com/sandronister/blockchain-go/internal/repository"
)

type BlockChain struct {
	LastHash []*entity.Block
	DB       repository.IRepository
}
