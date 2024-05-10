package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (b *BadgerRepository) GetBlock(hash []byte) (*entity.Block, error) {

	var block *entity.Block

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(hash))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			block = entity.Deserialize(val)
			return nil
		})

		return err
	})

	return block, err

}
