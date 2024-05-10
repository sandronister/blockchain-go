package badgerepository

import "github.com/dgraph-io/badger/v2"

func (b *BadgerRepository) GetLastHash() ([]byte, error) {
	var hash []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			hash = append([]byte{}, val...)
			return nil
		})
		return err
	})
	return hash, err
}
