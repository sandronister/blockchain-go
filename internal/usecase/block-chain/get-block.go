package blockchain

import "github.com/sandronister/blockchain-go/internal/entity"

func (b *BlockChain) GetBlock(hash []byte) (*entity.Block, error) {
	block, err := b.Repository.GetBlock(hash)

	if err != nil {
		return nil, err
	}

	return block, nil
}
