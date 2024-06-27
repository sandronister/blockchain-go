package repository

import "github.com/sandronister/blockchain-go/internal/entity"

type IRepository interface {
	GetBlock(hash []byte) (*entity.Block, error)
	GetLastHash() ([]byte, error)
	SetBlock(block entity.Block) error
	SetLastHash(block []byte) error
	InitBlock(entity.Block) error
	Close() error
	GetKeys() ([]string, error)
}
