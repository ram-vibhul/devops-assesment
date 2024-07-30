package db

import (
	"context"
	"database/sql"
	"simplebank/pkg/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomAmount(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, account.ID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	originalEntry := createRandomEntry(t)
	resultEntry, err := testQueries.GetEntry(context.Background(), originalEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, resultEntry)

	require.Equal(t, originalEntry.ID, resultEntry.ID)
	require.Equal(t, originalEntry.AccountID, resultEntry.AccountID)
	require.Equal(t, originalEntry.Amount, resultEntry.Amount)
	require.WithinDuration(t, originalEntry.CreatedAt, resultEntry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	originalEntry := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     originalEntry.ID,
		Amount: util.RandomAmount(),
	}
	resultEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, resultEntry)

	require.Equal(t, originalEntry.ID, resultEntry.ID)
	require.Equal(t, originalEntry.AccountID, resultEntry.AccountID)
	require.Equal(t, arg.Amount, resultEntry.Amount)
	require.WithinDuration(t, originalEntry.CreatedAt, resultEntry.CreatedAt, time.Second)

}

func TestDeleteEntry(t *testing.T) {
	originalEntry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), originalEntry.ID)
	require.NoError(t, err)

	resultEntry, err := testQueries.GetEntry(context.Background(), originalEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, resultEntry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
