package infra

import (
	"context"
	"database/sql"
	"fmt"
)

type TxAdmin struct {
	db *sql.DB
}

func NewTxAdmin(db *sql.DB) *TxAdmin {
	return &TxAdmin{db: db}
}

func (t *TxAdmin) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}
	return tx.Commit()
}
