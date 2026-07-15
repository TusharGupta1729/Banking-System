package handlers

import (
	"net/http"
	"strconv"

	"banking-system/models"
	"banking-system/services"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	service *services.LoanService
}

type RepayLoanRequest struct {
	Amount float64 `json:"amount"`
}

func NewLoanHandler(service *services.LoanService) *LoanHandler {
	return &LoanHandler{
		service: service,
	}
}

func (h *LoanHandler) CreateLoan(c *gin.Context) {

	var loan models.Loan

	if err := c.BindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	customerID, _ := c.Get("customer_id")

	loan.CustomerID = customerID.(uint)

	if err := h.service.CreateLoan(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func (h *LoanHandler) GetLoans(c *gin.Context) {

	role, _ := c.Get("role")

	if role.(string) == "admin" {

		loans, err := h.service.GetLoans()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, loans)
		return
	}

	customerID, _ := c.Get("customer_id")

	loans, err := h.service.GetLoansByCustomerID(customerID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func (h *LoanHandler) ApproveLoan(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid loan id",
		})
		return
	}

	err = h.service.ApproveLoan(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "loan approved successfully",
	})
}

func (h *LoanHandler) RejectLoan(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid loan id",
		})
		return
	}

	err = h.service.RejectLoan(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "loan rejected successfully",
	})
}

func (h *LoanHandler) RepayLoan(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid loan id",
		})
		return
	}

	var req RepayLoanRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	loan, err := h.service.GetLoanByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "loan not found",
		})
		return
	}

	customerID, _ := c.Get("customer_id")

	if loan.CustomerID != customerID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
		return
	}

	err = h.service.RepayLoan(
		uint(id),
		req.Amount,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "loan repayment successful",
	})
}
