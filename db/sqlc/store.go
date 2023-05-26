package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TxStore interface {
	Querier
	CreateCaseTx(context context.Context, args CreateCaseParams) (EcdeCase, error)
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) TxStore {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(context context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(context, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if err == tx.Rollback() {
			return fmt.Errorf(err.Error())
		}
		return err
	}
	return tx.Commit()
}
