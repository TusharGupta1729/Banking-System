package handlers

import (
	"net/http"

	"banking-system/models"
	"banking-system/services"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	service *services.LoanService
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

	if err := h.service.CreateLoan(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func (h *LoanHandler) GetLoans(c *gin.Context) {

	loans, err := h.service.GetLoans()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, loans)
}
