package routes

import (
	"banking-system/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, bankHandler *handlers.BankHandler) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Banking API Running",
		})
	})

	r.POST("/banks", bankHandler.CreateBank)
	r.GET("/banks", bankHandler.GetBanks)
}
