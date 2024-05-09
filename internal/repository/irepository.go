package repository

import "github.com/sandronister/blockchain-go/internal/entity"

type IRepository interface {
	UpdateBlockChain(block *entity.Block) error
	GetBlock(hash string) (*entity.Block, error)
	SetBlock(block *entity.Block) error
	SetLastHash(block []byte) error
}
