package blockchain

import "github.com/sandronister/blockchain-go/internal/entity"

func (bci *BlockChainIterator) Next() (*entity.Block, error) {
	block, err := bci.repository.GetBlock(bci.lastHash)
	if err != nil {
		return nil, err
	}
	bci.lastHash = block.Hash
	return block, nil
}
