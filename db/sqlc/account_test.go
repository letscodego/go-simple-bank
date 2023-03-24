package db

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/letscodego/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	lastId, err := account.LastInsertId()

	// get the account we just inserted
	fetchedAccount, err := testQueries.GetAccount(context.Background(), lastId)
	if err != nil {
		log.Fatal("cannot retrieve last inserted row:", err)
	}

	require.Equal(t, arg.Owner, fetchedAccount.Owner)
	require.Equal(t, arg.Balance, fetchedAccount.Balance)
	require.Equal(t, arg.Currency, fetchedAccount.Currency)

	require.NotZero(t, fetchedAccount.ID)
	require.NotZero(t, fetchedAccount.CreatedAt)

	return fetchedAccount
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{Balance: util.RandomMoney(), ID: account1.ID}
	err := testQueries.UpdateAccount(context.Background(), arg)

	// get the account we just updated
	account2, e := testQueries.GetAccount(context.Background(), account1.ID)
	if e != nil {
		log.Fatal("cannot retrieve last inserted row:", err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Offset: 5,
		Limit:  5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
