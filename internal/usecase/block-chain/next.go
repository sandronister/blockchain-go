package blockchain

import (
	"fmt"

	"github.com/sandronister/blockchain-go/internal/entity"
)

func (bci *BlockChainIterator) Next() (*entity.Block, error) {

	if bci.lastHash == nil {
		return nil, nil
	}

	block, err := bci.repository.GetBlock(bci.lastHash)
	if err != nil {
		return nil, err
	}

	fmt.Println("Block data", block.Data)
	bci.lastHash = block.PrevHash
	return block, nil
}
