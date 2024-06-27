package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (bc *BlockChain) AddBlock(data string) error {
	block := entity.CreateBlock(data, bc.LastHash)
	err := bc.Repository.SetBlock(block)

	if err != nil {
		return err
	}

	bc.LastHash = block.Hash
	err = bc.Repository.SetLastHash(block.Hash)

	if err != nil {
		return err
	}

	return nil
}
