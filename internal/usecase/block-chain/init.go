package blockchain

import (
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (bc *BlockChain) Init() error {

	lastHash, err := bc.Repository.GetLastHash()

	if err != nil {
		block := entity.Genesis()
		err = bc.Repository.InitBlock(block)
		if err != nil {
			return err
		}

		lastHash = block.Hash

	}

	bc.LastHash = lastHash

	return nil
}
