package routes

import (
	"banking-system/handlers"
	"banking-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	bankHandler *handlers.BankHandler,
	branchHandler *handlers.BranchHandler,
	customerHandler *handlers.CustomerHandler,
	accountHandler *handlers.AccountHandler,
	transactionHandler *handlers.TransactionHandler,
	loanHandler *handlers.LoanHandler,
	authHandler *handlers.AuthHandler,
) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Banking API Running",
			"status":  "Server is running successfully",
			"usage":   "Use Postman or any API client to access the available endpoints",
		})
	})

	r.POST("/banks", bankHandler.CreateBank)
	r.GET("/banks", bankHandler.GetBanks)

	r.POST("/branches", branchHandler.CreateBranch)
	r.GET("/branches", branchHandler.GetBranches)

	r.POST("/customers", customerHandler.CreateCustomer)
	r.GET("/customers", customerHandler.GetCustomers)
	r.GET("/customers/:id/accounts", customerHandler.GetCustomerAccounts)

	r.POST("/accounts", accountHandler.CreateAccount)
	r.GET("/accounts", accountHandler.GetAccounts)
	r.POST(
		"/accounts/:id/deposit",
		middleware.AuthMiddleware(),
		accountHandler.Deposit,
	)
	r.POST(
		"/accounts/:id/withdraw",
		middleware.AuthMiddleware(),
		accountHandler.Withdraw,
	)
	r.GET("/accounts/:id", accountHandler.GetAccountByID)
	r.POST(
		"/accounts/transfer",
		middleware.AuthMiddleware(),
		accountHandler.Transfer,
	)

	r.POST("/transactions", transactionHandler.CreateTransaction)
	r.GET("/transactions", transactionHandler.GetTransactions)
	r.GET("/accounts/:id/transactions", accountHandler.GetAccountTransactions)

	r.POST("/loans", loanHandler.CreateLoan)
	r.GET("/loans", loanHandler.GetLoans)
	r.POST("/loans/:id/approve", loanHandler.ApproveLoan)
	r.POST("/loans/:id/reject", loanHandler.RejectLoan)
	r.POST("/loans/:id/repay", loanHandler.RepayLoan)

	r.POST("/login", authHandler.Login)
}
