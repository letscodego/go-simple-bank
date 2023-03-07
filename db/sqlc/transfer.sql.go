// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: transfer.sql

package db

import (
	"context"
	"database/sql"
)

const createTransfer = `-- name: CreateTransfer :execresult
INSERT INTO transfers (
  from_account_id, 
  to_account_id,
  amount
) VALUES (
  ?, ?, ? )
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = ?
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT ?, ?
`

type ListTransfersParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :execresult
UPDATE transfers 
SET amount = ?
WHERE id = ?
`

type UpdateTransferParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTransfer, arg.Amount, arg.ID)
}
