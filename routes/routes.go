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

	r.POST(
		"/banks",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		bankHandler.CreateBank,
	)
	r.GET("/banks", bankHandler.GetBanks)

	r.POST(
		"/branches",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		branchHandler.CreateBranch,
	)
	r.GET("/branches", branchHandler.GetBranches)

	r.POST("/customers", customerHandler.CreateCustomer)
	r.GET(
		"/customers",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		customerHandler.GetCustomers,
	)
	r.GET(
		"/customers/:id/accounts",
		middleware.AuthMiddleware(),
		customerHandler.GetCustomerAccounts,
	)

	r.POST(
		"/accounts",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		accountHandler.CreateAccount,
	)

	r.GET(
		"/accounts",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		accountHandler.GetAccounts,
	)
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
	r.GET(
		"/accounts/:id",
		middleware.AuthMiddleware(),
		accountHandler.GetAccountByID,
	)
	r.POST(
		"/accounts/transfer",
		middleware.AuthMiddleware(),
		accountHandler.Transfer,
	)
	r.GET(
		"/accounts/:id/transactions",
		middleware.AuthMiddleware(),
		accountHandler.GetAccountTransactions,
	)

	r.GET(
		"/transactions",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		transactionHandler.GetTransactions,
	)

	r.POST(
		"/loans",
		middleware.AuthMiddleware(),
		loanHandler.CreateLoan,
	)

	r.GET(
		"/loans",
		middleware.AuthMiddleware(),
		loanHandler.GetLoans,
	)

	r.POST(
		"/loans/:id/approve",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		loanHandler.ApproveLoan,
	)

	r.POST(
		"/loans/:id/reject",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		loanHandler.RejectLoan,
	)

	r.POST(
		"/loans/:id/repay",
		middleware.AuthMiddleware(),
		loanHandler.RepayLoan,
	)

	r.POST("/login", authHandler.Login)
}
