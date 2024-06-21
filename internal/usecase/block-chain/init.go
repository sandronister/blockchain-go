package blockchain

import (
	"fmt"

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
		fmt.Println(lastHash, "---------------------------------------------")

	}

	fmt.Printf("Last hash: %x\n ", lastHash)

	bc.LastHash = lastHash

	return nil
}
