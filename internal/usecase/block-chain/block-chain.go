package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/entity"
	"github.com/sandronister/blockchain-go/internal/repository"
)

type BlockChain struct {
	lashHash   []byte
	repository repository.IRepository
}

func NewBlockChain(repository repository.IRepository) *BlockChain {
	return &BlockChain{repository: repository}
}

func (bc *BlockChain) AddBlock(data string) error {
	block := entity.CreateBlock(data, bc.lashHash)
	err := bc.repository.Update(block)

	if err != nil {
		return err
	}

	bc.lashHash = block.Hash

	return nil
}
