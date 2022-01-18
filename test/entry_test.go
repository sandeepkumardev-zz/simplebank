package test

import (
	"context"
	"testing"
	"time"

	models "github.com/sandeepkumardev/simplebank/models"
	"github.com/sandeepkumardev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func getAccountId(t *testing.T) []int64 {
	arg := models.ListAccountsParams{
		Limit:  2,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 2)
	var accountId []int64
	for _, account := range accounts {
		accountId = append(accountId, account.ID)
	}

	return accountId
}

func createRandomEntry(t *testing.T) models.Entry {
	arg := models.CreateEntryParams{
		AccountID: getAccountId(t)[0],
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.AccountID, entry1.AccountID)
	require.Equal(t, entry2.Amount, entry1.Amount)

	require.NotZero(t, entry2.ID)
	require.WithinDuration(t, entry2.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := models.ListEntriesParams{
		AccountID: getAccountId(t)[0],
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
