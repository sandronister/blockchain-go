package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
	"github.com/sandronister/blockchain-go/pkg/catch"
)

func (br *BadgerRepository) Update(block *entity.Block) error {
	err := br.db.Update(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte("lh"))
		catch.Handle(err)

		err = item.Value(func(val []byte) error {

			err = txn.Set([]byte("lh"), val)
			catch.Handle(err)

			return nil
		})

		return txn.Set([]byte(block.Hash), block.Serialize())

	})

	if err != nil {
		return err
	}

	return br.SetLastHash(block.Serialize())
}
