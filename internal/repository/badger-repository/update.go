package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
	"github.com/sandronister/blockchain-go/pkg/catch"
)

func (br *BadgerRepository) UpdateBlockChain(block *entity.Block) error {
	err := br.db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err != badger.ErrKeyNotFound {
			genesis := entity.Genesis()
			err := txn.Set([]byte(genesis.Hash), genesis.Serialize())

			catch.Handle(err)
		}

		return txn.Set([]byte(block.Hash), block.Serialize())

	})

	if err != nil {
		return err
	}

	return br.SetLastHash(block.Serialize())
}
