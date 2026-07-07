package handlers

import (
	"banking-system/models"
	"banking-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankHandler struct {
	service *services.BankService
}

func NewBankHandler(service *services.BankService) *BankHandler {
	return &BankHandler{
		service: service,
	}
}

func (h *BankHandler) CreateBank(c *gin.Context) {
	var bank models.Bank

	if err := c.BindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.service.CreateBank(&bank)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, bank)
}

func (h *BankHandler) GetBanks(c *gin.Context) {

	banks, err := h.service.GetBanks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, banks)
}
