package repository

import "github.com/sandronister/blockchain-go/internal/entity"

type IRepository interface {
	Update(block *entity.Block) error
	GetBlock(hash []byte) (*entity.Block, error)
	GetLastHash() ([]byte, error)
	SetBlock(block *entity.Block) error
	SetLastHash(block []byte) error
	InitBlock() (*entity.Block, error)
}
