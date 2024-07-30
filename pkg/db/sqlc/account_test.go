package db

import (
	"context"
	"database/sql"
	"simplebank/pkg/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, arg.Owner)
	require.Equal(t, arg.Balance, arg.Balance)
	require.Equal(t, arg.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountOriginal := createRandomAccount(t)
	accountResult, err := testQueries.GetAccount(context.Background(), accountOriginal.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accountResult)

	require.Equal(t, accountOriginal.ID, accountResult.ID)
	require.Equal(t, accountOriginal.Owner, accountResult.Owner)
	require.Equal(t, accountOriginal.Balance, accountResult.Balance)
	require.Equal(t, accountOriginal.Currency, accountResult.Currency)
	require.WithinDuration(t, accountOriginal.CreatedAt.Time, accountResult.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accountOriginal := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      accountOriginal.ID,
		Balance: util.RandomMoney(),
	}
	accountResult, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountResult)

	require.Equal(t, accountOriginal.ID, accountResult.ID)
	require.Equal(t, accountOriginal.Owner, accountResult.Owner)
	require.Equal(t, arg.Balance, accountResult.Balance)
	require.Equal(t, accountOriginal.Currency, accountResult.Currency)
	require.WithinDuration(t, accountOriginal.CreatedAt.Time, accountResult.CreatedAt.Time, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	accountOriginal := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accountOriginal.ID)
	require.NoError(t, err)

	accountResult, err := testQueries.GetAccount(context.Background(), accountOriginal.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountResult)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accouts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accouts, 5)

	for _, account := range accouts {
		require.NotEmpty(t, account)
	}
}
