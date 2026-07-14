package handlers

import (
	"net/http"
	"strconv"

	"banking-system/models"
	"banking-system/services"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service *services.AccountService
}

type DepositRequest struct {
	Amount float64 `json:"amount"`
}

type WithdrawRequest struct {
	Amount float64 `json:"amount"`
}

type TransferRequest struct {
	FromAccount uint    `json:"fromAccount"`
	ToAccount   uint    `json:"toAccount"`
	Amount      float64 `json:"amount"`
}

func NewAccountHandler(service *services.AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {

	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := h.service.CreateAccount(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (h *AccountHandler) GetAccounts(c *gin.Context) {

	accounts, err := h.service.GetAccounts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (h *AccountHandler) Deposit(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid account id",
		})
		return
	}

	var req DepositRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = h.service.Deposit(uint(id), req.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deposit successful",
	})
}

func (h *AccountHandler) Withdraw(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid account id",
		})
		return
	}

	var req WithdrawRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	account, err := h.service.GetAccountByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "account not found",
		})
		return
	}

	customerID, _ := c.Get("customer_id")

	if account.CustomerID != customerID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
		return
	}

	err = h.service.Withdraw(uint(id), req.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "withdraw successful",
	})
}

func (h *AccountHandler) GetAccountByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid account id",
		})
		return
	}

	account, err := h.service.GetAccountByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "account not found",
		})
		return
	}

	role, _ := c.Get("role")

	if role != "admin" {

		customerID, _ := c.Get("customer_id")

		if account.CustomerID != customerID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "access denied",
			})
			return
		}
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) Transfer(c *gin.Context) {

	var req TransferRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	fromAccount, err := h.service.GetAccountByID(
		req.FromAccount,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "source account not found",
		})
		return
	}

	customerID, _ := c.Get("customer_id")

	if fromAccount.CustomerID != customerID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
		return
	}

	err = h.service.Transfer(
		req.FromAccount,
		req.ToAccount,
		req.Amount,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transfer successful",
	})
}

func (h *AccountHandler) GetAccountTransactions(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid account id",
		})
		return
	}
	account, err := h.service.GetAccountByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "account not found",
		})
		return
	}

	role, _ := c.Get("role")

	if role.(string) != "admin" {

		customerID, _ := c.Get("customer_id")

		if account.CustomerID != customerID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "access denied",
			})
			return
		}
	}

	transactions, err := h.service.GetAccountTransactions(
		uint(id),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
