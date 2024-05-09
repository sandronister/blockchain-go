package badgerepository

import "github.com/dgraph-io/badger/v2"

func (b *BadgerRepository) SetLastHash(hash []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("lh"), hash)
		return err
	})
}
