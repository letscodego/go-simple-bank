package db

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/letscodego/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// get the account we just inserted
	fetchedUser, err := testQueries.GetUser(context.Background(), arg.Username)
	if err != nil {
		log.Fatal("cannot retrieve last inserted row:", err)
	}

	require.Equal(t, arg.Email, fetchedUser.Email)
	require.Equal(t, arg.FullName, fetchedUser.FullName)
	require.Equal(t, arg.HashedPassword, fetchedUser.HashedPassword)
	require.Equal(t, arg.Username, fetchedUser.Username)

	//require.True(t, fetchedUser.PasswordChangedAt.IsZero())
	require.NotZero(t, fetchedUser.CreatedAt)

	return fetchedUser
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Username, user2.Username)
	//require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
