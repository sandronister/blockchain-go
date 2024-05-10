package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/internal/entity"
)

func (br *BadgerRepository) InitBlock() (*entity.Block, error) {
	genesisBlock := entity.Genesis()
	err := br.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(genesisBlock.Hash), genesisBlock.Serialize())
		if err != nil {
			return err
		}
		return txn.Set([]byte("lh"), []byte(genesisBlock.Hash))
	})

	return genesisBlock, err
}
