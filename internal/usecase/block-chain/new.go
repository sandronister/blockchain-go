package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/repository"
)

type BlockChain struct {
	LastHash   []byte
	Repository repository.IRepository
}

type BlockChainIterator struct {
	lastHash   []byte
	repository repository.IRepository
}

func NewBlockChain(repository repository.IRepository) *BlockChain {
	return &BlockChain{Repository: repository, LastHash: nil}
}
