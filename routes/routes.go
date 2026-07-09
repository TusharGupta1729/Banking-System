package routes

import (
	"banking-system/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	bankHandler *handlers.BankHandler,
	branchHandler *handlers.BranchHandler,
	customerHandler *handlers.CustomerHandler,
	accountHandler *handlers.AccountHandler,
	transactionHandler *handlers.TransactionHandler,
	loanHandler *handlers.LoanHandler,
) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Banking API Running",
		})
	})

	r.POST("/banks", bankHandler.CreateBank)
	r.GET("/banks", bankHandler.GetBanks)

	r.POST("/branches", branchHandler.CreateBranch)
	r.GET("/branches", branchHandler.GetBranches)

	r.POST("/customers", customerHandler.CreateCustomer)
	r.GET("/customers", customerHandler.GetCustomers)

	r.POST("/accounts", accountHandler.CreateAccount)
	r.GET("/accounts", accountHandler.GetAccounts)

	r.POST("/transactions", transactionHandler.CreateTransaction)
	r.GET("/transactions", transactionHandler.GetTransactions)

	r.POST("/loans", loanHandler.CreateLoan)
	r.GET("/loans", loanHandler.GetLoans)
}
