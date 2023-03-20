// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (sql.Result, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (sql.Result, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) error
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) error
}

var _ Querier = (*Queries)(nil)
