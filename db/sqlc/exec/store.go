package db

import (
	"context"
	"database/sql"
	"fmt"

	sqlc "github.com/hidratechh/GO_WMS/db/sqlc"
	txData "github.com/hidratechh/GO_WMS/db/sqlc/dto"
)

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*sqlc.Queries
	db *sql.DB
}

//create a new store
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

//execute function within a db trans
func (store *SQLStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

/*
===EXTENDED sqlc.Queries===
query with custom bussiness logic, builded manually without sqlc
*/
func (store *SQLStore) TransferTx(arg txData.TransferTxParams) {

}
