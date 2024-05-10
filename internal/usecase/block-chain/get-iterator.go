package blockchain

func (bc *BlockChain) GetIterator() *BlockChainIterator {
	return &BlockChainIterator{
		currentHash: bc.LastHash,
		repository:  bc.Repository,
	}
}
