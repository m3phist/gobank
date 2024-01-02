package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/m3phist/gobank/backend/db/sqlc"
	"github.com/m3phist/gobank/backend/utils"
)

type Account struct {
	server *Server
}

func (a Account) router(server *Server) {
	a.server = server

	serverGroup := server.router.Group("/account", AuthenticatedMiddleware())
	serverGroup.POST("create", a.createAccount)
	serverGroup.GET("", a.getUserAccounts)
	serverGroup.POST("transfer", a.transfer)
}

type AccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (a *Account) createAccount(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	acc := new(AccountRequest)

	if err := c.ShouldBindJSON(acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateAccountParams{
		UserID:   int32(userId),
		Currency: acc.Currency,
		Balance:  0,
	}

	account, err := a.server.queries.CreateAccount(context.Background(), arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "you already have an account with this currency"})
				return
			}

			if pgErr.Code == "23503" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
				return
			}

		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, account)

}

func (a *Account) getUserAccounts(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	accounts, err := a.server.queries.GetAccountByUserID(context.Background(), int32(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

type TransferRequest struct {
	ToAccountID   int32   `json:"to_account_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	FromAccountID int32   `json:"from_account_id" binding:"required"`
}

func (a *Account) transfer(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	tr := new(TransferRequest)

	if err := c.ShouldBindJSON(&tr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := a.server.queries.GetAccountByID(context.Background(), int64(tr.FromAccountID))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to fetch the account"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if account.UserID != int32(userId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to fetch the account"})
		return
	}

	toAccount, err := a.server.queries.GetAccountByID(context.Background(), int64(tr.ToAccountID))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find the account to send the request"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if toAccount.Currency != account.Currency {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transfer currency does not match with recipient account"})
		return
	}

	if account.Balance < tr.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	txArg := db.CreateTransferParams{
		FromAccountID: tr.FromAccountID,
		ToAccountID:   tr.ToAccountID,
		Amount:        tr.Amount,
	}

	tx, err := a.server.queries.TransferTx(context.Background(), txArg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encountered an issue with transaction"})
		return
	}

	c.JSON(http.StatusCreated, tx)
}
