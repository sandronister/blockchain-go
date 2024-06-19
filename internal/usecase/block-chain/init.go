package blockchain

func (bc *BlockChain) Init() error {
	block, err := bc.Repository.InitBlock()
	if err != nil {
		return err
	}
	bc.LastHash = block.Hash
	return nil
}
