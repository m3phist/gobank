package db_test

import (
	"context"
	"log"
	"testing"
	"time"

	db "github.com/m3phist/gobank/backend/db/sqlc"
	"github.com/m3phist/gobank/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	hashedPassword, err := utils.GenerateHashPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Unable to generate hash password", err)
	}

	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, user.Email, arg.Email)
	assert.Equal(t, user.HashedPassword, arg.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

	user2, err := testQuery.CreateUser(context.Background(), arg)
	assert.Error(t, err)
	assert.Empty(t, user2)
}
