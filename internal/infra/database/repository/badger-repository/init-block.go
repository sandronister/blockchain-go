package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (br *BadgerRepository) InitBlock(block entity.Block) error {

	err := br.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(block.Hash), block.Serialize())
		if err != nil {
			return err
		}
		return txn.Set([]byte("lh"), []byte(block.Hash))
	})

	return err
}
