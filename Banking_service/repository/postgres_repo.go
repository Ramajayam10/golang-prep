package repository

import (
	"banking-service/models"

	"gorm.io/gorm"
)

type PostgresBankRepository struct {
	db *gorm.DB
}

func NewPostgresBankRepository(db *gorm.DB) BankRepository {
	return &PostgresBankRepository{db: db}
}

// 1. User operations
func (r *PostgresBankRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *PostgresBankRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresBankRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 2. User Credentials operations
func (r *PostgresBankRepository) CreateCredential(cred *models.UserCredential) error {
	return r.db.Create(cred).Error
}

func (r *PostgresBankRepository) GetCredentialByUserID(userID uint) (*models.UserCredential, error) {
	var cred models.UserCredential
	err := r.db.Where("user_id = ?", userID).First(&cred).Error
	if err != nil {
		return nil, err
	}
	return &cred, nil
}

// 3. Account operations
func (r *PostgresBankRepository) CreateAccount(account *models.BankAccount) error {
	return r.db.Create(account).Error
}

func (r *PostgresBankRepository) GetAccountByNumber(accountNumber string) (*models.BankAccount, error) {
	var account models.BankAccount
	err := r.db.Where("account_number = ?", accountNumber).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *PostgresBankRepository) GetAccountByID(id uint) (*models.BankAccount, error) {
	var account models.BankAccount
	err := r.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// 4. Financial Transactions
func (r *PostgresBankRepository) TransferFunds(senderAccountID, receiverAccountID uint, amount float64, txn *models.Transaction) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Step A: Deduct from Sender
		if err := tx.Model(&models.BankAccount{}).
			Where("id = ?", senderAccountID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		// Step B: Add to Receiver
		if err := tx.Model(&models.BankAccount{}).
			Where("id = ?", receiverAccountID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// Step C: Save Transaction
		if err := tx.Create(txn).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *PostgresBankRepository) GetTransactionsByAccount(accountID uint) ([]models.Transaction, error) {
	var txns []models.Transaction
	err := r.db.Where("sender_account_id = ? OR receiver_account_id = ?", accountID, accountID).
		Order("created_at desc").
		Find(&txns).Error
	return txns, err
}
