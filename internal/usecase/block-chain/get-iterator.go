package blockchain

func (bc *BlockChain) GetIterator() *BlockChainIterator {
	return &BlockChainIterator{
		lastHash:   bc.LastHash,
		repository: bc.Repository,
	}
}
