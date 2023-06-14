package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxStore interface {
	Querier
	CreateCaseTx(context context.Context, args CreateCaseParams) (EcdeCase, error)
}

type Store struct {
	*Queries
	pool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) TxStore {
	return &Store{
		pool:    connPool,
		Queries: New(connPool),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return tx.Rollback(ctx)
		}
		return err
	}
	return tx.Commit(ctx)
}
