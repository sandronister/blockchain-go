package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (b *BadgerRepository) SetBlock(block *entity.Block) error {

	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(block.Hash, block.Serialize())
		return err
	})
}
