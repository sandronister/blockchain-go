package badger

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/sandronister/blockchain-go/configs"
)

func GetConnection(cfg *configs.Conf) (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions(cfg.DBPath))
	if err != nil {
		return nil, err
	}

	return db, nil
}
