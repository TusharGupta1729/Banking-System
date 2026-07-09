package handlers

import (
	"net/http"

	"banking-system/models"
	"banking-system/services"

	"github.com/gin-gonic/gin"
)

type BranchHandler struct {
	service *services.BranchService
}

func NewBranchHandler(service *services.BranchService) *BranchHandler {
	return &BranchHandler{
		service: service,
	}
}

func (h *BranchHandler) CreateBranch(c *gin.Context) {

	var branch models.Branch

	if err := c.BindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := h.service.CreateBranch(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

func (h *BranchHandler) GetBranches(c *gin.Context) {

	branches, err := h.service.GetBranches()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, branches)
}
