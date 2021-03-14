package handler

import (
	"bwastartup/transaction"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) Index(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransactions()
	if err != nil {
		c.HTML(500, "error.html", nil)
		return
	}
	c.HTML(200, "transaction_index.html", gin.H{"transactions": transactions})
}
