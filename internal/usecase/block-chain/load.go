package blockchain

import (
	"fmt"

	"github.com/sandronister/blockchain-go/internal/entity"
)

func (b *BlockChain) Load() error {
	block, err := b.Repository.GetLastHash()

	if err != nil {
		return err
	}

	b.LastHash = entity.Deserialize(block).Hash

	for {
		block, err := b.Repository.GetBlock(b.LastHash)

		if err != nil {
			return err
		}

		fmt.Printf("Loaded block: %x\n", block.Hash)

		if block.PrevHash == nil {
			break
		}

		b.LastHash = block.PrevHash
	}

	return nil
}
