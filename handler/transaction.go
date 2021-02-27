package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", 400, "Failed", nil)
		c.JSON(400, response)
		return
	}

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", 400, "Failed", nil)
		c.JSON(400, response)
		return
	}
	response := helper.APIResponse("Campaign's transactions", 200, "Success", transactions)
	c.JSON(200, response)
}
