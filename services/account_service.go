package services

import (
	"errors"

	"banking-system/config"
	"banking-system/models"
	"banking-system/repository"
)

type AccountService struct {
	repo            *repository.AccountRepository
	transactionRepo *repository.TransactionRepository
}

type TransferRequest struct {
	FromAccount uint    `json:"fromAccount"`
	ToAccount   uint    `json:"toAccount"`
	Amount      float64 `json:"amount"`
}

func NewAccountService(
	repo *repository.AccountRepository,
	transactionRepo *repository.TransactionRepository,
) *AccountService {

	return &AccountService{
		repo:            repo,
		transactionRepo: transactionRepo,
	}
}

func (s *AccountService) CreateAccount(account *models.Account) error {

	if account.AccountNumber == "" {
		return errors.New("account number is required")
	}

	if account.CustomerID == 0 {
		return errors.New("customer id is required")
	}

	if account.BranchID == 0 {
		return errors.New("branch id is required")
	}

	return s.repo.Create(account)
}

func (s *AccountService) GetAccounts() ([]models.Account, error) {
	return s.repo.GetAll()
}

func (s *AccountService) Deposit(
	accountID uint,
	amount float64,
) error {

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	account, err := s.repo.GetByID(accountID)

	if err != nil {
		return err
	}

	account.Balance += amount

	if err := s.repo.Update(account); err != nil {
		return err
	}

	transaction := &models.Transaction{
		AccountID: account.ID,
		Type:      "Deposit",
		Amount:    amount,
	}

	return s.transactionRepo.Create(transaction)
}

func (s *AccountService) Withdraw(
	accountID uint,
	amount float64,
) error {

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	account, err := s.repo.GetByID(accountID)

	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount

	if err := s.repo.Update(account); err != nil {
		return err
	}

	transaction := &models.Transaction{
		AccountID: account.ID,
		Type:      "Withdraw",
		Amount:    amount,
	}

	return s.transactionRepo.Create(transaction)
}

func (s *AccountService) GetAccountByID(id uint) (*models.Account, error) {
	return s.repo.GetByID(id)
}

func (s *AccountService) Transfer(
	fromID uint,
	toID uint,
	amount float64,
) error {

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if fromID == toID {
		return errors.New("cannot transfer to same account")
	}

	fromAccount, err := s.repo.GetByID(fromID)
	if err != nil {
		return errors.New("sender account not found")
	}

	toAccount, err := s.repo.GetByID(toID)
	if err != nil {
		return errors.New("receiver account not found")
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient balance")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	tx := config.DB.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	if err := s.repo.UpdateWithTransaction(tx, fromAccount); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.repo.UpdateWithTransaction(tx, toAccount); err != nil {
		tx.Rollback()
		return err
	}

	transferOut := &models.Transaction{
		AccountID: fromAccount.ID,
		Type:      "Transfer Out",
		Amount:    amount,
	}

	if err := s.transactionRepo.CreateWithTransaction(tx, transferOut); err != nil {
		tx.Rollback()
		return err
	}

	transferIn := &models.Transaction{
		AccountID: toAccount.ID,
		Type:      "Transfer In",
		Amount:    amount,
	}

	if err := s.transactionRepo.CreateWithTransaction(tx, transferIn); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *AccountService) GetAccountTransactions(
	id uint,
) ([]models.Transaction, error) {

	account, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return account.Transactions, nil
}
