package badgerepository

import "github.com/dgraph-io/badger/v2"

func (r *BadgerRepository) GetKeys() ([]string, error) {
	keys := make([]string, 0)
	err := r.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			keys = append(keys, string(k))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return keys, nil
}
