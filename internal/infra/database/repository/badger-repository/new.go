package badgerepository

import (
	"github.com/dgraph-io/badger/v2"
)

type BadgerRepository struct {
	db *badger.DB
}

func NewBadgerRepository(db *badger.DB) *BadgerRepository {
	return &BadgerRepository{db: db}
}
