package db

import "context"

func (s *Store) CreateCaseTx(context context.Context, args CreateCaseParams) (EcdeCase, error) {
	var ecde EcdeCase
	var err error
	err = s.execTx(context, func(q *Queries) error {
		ecde, err = q.CreateCase(context, args)

		if err != nil {
			return err
		}

		return nil
	})
	return ecde, err
}
