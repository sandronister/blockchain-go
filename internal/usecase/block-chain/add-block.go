package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (bc *BlockChain) AddBlock(data string) error {
	block := entity.CreateBlock(data, bc.LastHash)
	err := bc.Repository.Update(block)

	if err != nil {
		return err
	}

	bc.LastHash = block.Hash

	return nil
}
