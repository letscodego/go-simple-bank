package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	lastId, err := account.LastInsertId()

	// get the account we just inserted
	fetchedAccount, err := testQueries.GetAccount(context.Background(), lastId)
	if err != nil {
		log.Fatal("cannot retrive last inserted row:", err)
	}

	require.Equal(t, arg.Owner, fetchedAccount.Owner)
	require.Equal(t, arg.Balance, fetchedAccount.Balance)
	require.Equal(t, arg.Currency, fetchedAccount.Currency)

	require.NotZero(t, fetchedAccount.ID)
	require.NotZero(t, fetchedAccount.CreatedAt)
}
