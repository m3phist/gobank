package db_test

import (
	"context"
	"testing"
	"time"

	db "github.com/m3phist/gobank/backend/db/sqlc"
	"github.com/stretchr/testify/assert"
)

func createRandomAccount(user_id int64, t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		UserID:   int32(user_id),
		Currency: "USD",
		Balance:  500,
	}
	account, err := testQuery.CreateAccount(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, account)

	assert.Equal(t, account.UserID, arg.UserID)
	assert.Equal(t, account.Currency, arg.Currency)
	assert.Equal(t, account.Balance, arg.Balance)
	assert.WithinDuration(t, account.CreatedAt, time.Now(), 2*time.Second)

	return account
}

func TestTransfer(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	account1 := createRandomAccount(user1.ID, t)
	account2 := createRandomAccount(user2.ID, t)

	arg := db.CreateTransferParams{
		FromAccountID: int32(account1.ID),
		ToAccountID:   int32(account2.ID),
		Amount:        200,
	}

	txResponseChan := make(chan db.TransferTxResponse)
	errChan := make(chan error)
	count := 10

	for i := 0; i < count; i++ {
		go func() {
			tx, err := testQuery.TransferTx(context.Background(), arg)
			errChan <- err
			txResponseChan <- tx
		}()
	}

	for i := 0; i < count; i++ {
		err := <-errChan
		tx := <-txResponseChan

		assert.NoError(t, err)
		assert.NotEmpty(t, tx)

		// test transfer
		assert.Equal(t, tx.Transfer.FromAccountID, arg.FromAccountID)
		assert.Equal(t, tx.Transfer.ToAccountID, arg.ToAccountID)
		assert.Equal(t, tx.Transfer.Amount, arg.Amount)

		// test entry
		// entry In
		assert.Equal(t, tx.EntryIn.AccountID, arg.ToAccountID)
		assert.Equal(t, tx.EntryIn.Amount, arg.Amount)

		// entry Out
		assert.Equal(t, tx.EntryOut.AccountID, arg.FromAccountID)
		assert.Equal(t, tx.EntryOut.Amount, -1*arg.Amount)

	}

	newAccount1, err := testQuery.GetAccountByID(context.Background(), account1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, newAccount1)

	newAccount2, err := testQuery.GetAccountByID(context.Background(), account2.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, newAccount2)

	newAccount := float64(count * int(arg.Amount))
	assert.Equal(t, newAccount1.Balance, (account1.Balance - newAccount))
	assert.Equal(t, newAccount2.Balance, (account1.Balance + newAccount))
}
