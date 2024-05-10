package blockchain

import "github.com/sandronister/blockchain-go/internal/entity"

func (bci *BlockChainIterator) Next() (*entity.Block, error) {
	block, err := bci.repository.GetBlock(bci.currentHash)
	if err != nil {
		return nil, err
	}
	bci.currentHash = block.PrevHash
	return block, nil
}
