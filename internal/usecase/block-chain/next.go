package blockchain

import (
	"errors"

	"github.com/sandronister/blockchain-go/internal/entity"
)

func (bci *BlockChainIterator) Next() (*entity.Block, error) {

	if bci.lastHash == nil {
		return nil, errors.New("no more blocks")
	}

	block, err := bci.repository.GetBlock(bci.lastHash)
	if err != nil {
		return nil, err
	}

	bci.lastHash = block.PrevHash
	return block, nil
}
