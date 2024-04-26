package entity

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
