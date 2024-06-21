package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (br *BadgerRepository) Update(block entity.Block) error {
	err := br.db.Update(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte("lh"))

		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			return txn.Set([]byte("lh"), val)
		})

		if err != nil {
			return err
		}

		return txn.Set([]byte(block.Hash), block.Serialize())

	})

	if err != nil {
		return err
	}

	return nil
	//return br.SetLastHash(block.Serialize())
}
