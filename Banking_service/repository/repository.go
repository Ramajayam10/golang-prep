package repository

import "banking-service/models"

type BankRepository interface {
	// 1. User operations
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	// 2. User Credentials operations
	CreateCredential(cred *models.UserCredential) error
	GetCredentialByUserID(userID uint) (*models.UserCredential, error)

	// 3. Account operations
	CreateAccount(account *models.BankAccount) error
	GetAccountByNumber(accountNumber string) (*models.BankAccount, error)
	GetAccountByID(id uint) (*models.BankAccount, error)

	// 4. Financial Transactions
	TransferFunds(senderAccountID, receiverAccountID uint, amount float64, txn *models.Transaction) error
	GetTransactionsByAccount(accountID uint) ([]models.Transaction, error)
}
