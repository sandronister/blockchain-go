package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/repository"
)

type BlockChain struct {
	LastHash   []byte
	Repository repository.IRepository
}

type BlockChainIterator struct {
	currentHash []byte
	repository  repository.IRepository
}

func NewBlockChain(repository repository.IRepository) (*BlockChain, error) {
	block, err := repository.InitBlock()
	if err != nil {
		return nil, err
	}
	return &BlockChain{Repository: repository, LastHash: block.Hash}, nil
}
