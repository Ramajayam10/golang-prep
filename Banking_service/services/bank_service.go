package services

import (
	"fmt"
	"sync"

	"banking-service/models"
	"banking-service/repository"
	"banking-service/utils"
)

type BankService struct {
	repo repository.BankRepository
	mu   sync.Mutex
}

// NewBankService is our Constructor (Dependency Injection)
func NewBankService(repo repository.BankRepository) *BankService {
	return &BankService{
		repo: repo,
	}
}

// 1. Create a User (Registration with Password Hashing)
func (s *BankService) CreateUser(name, phone, email, address, state, country, password string, age int) (*models.User, error) {
	user := &models.User{
		Name:        name,
		PhoneNumber: phone,
		Age:         age,
		Email:       email,
		Address:     address,
		State:       state,
		Country:     country,
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	if password != "" {
		hashed, err := utils.HashPassword(password)
		if err != nil {
			return nil, err
		}
		cred := &models.UserCredential{
			UserID:   user.ID,
			Password: hashed,
		}
		if err := s.repo.CreateCredential(cred); err != nil {
			return nil, err
		}
	}

	return user, nil
}

// 2. User Login: Authenticate and issue JWT token
func (s *BankService) Login(email, password string) (string, *models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", nil, utils.ErrInvalidCredentials
	}

	cred, err := s.repo.GetCredentialByUserID(user.ID)
	if err != nil {
		return "", nil, utils.ErrInvalidCredentials
	}

	if !utils.CheckPassword(cred.Password, password) {
		return "", nil, utils.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// 3. Open a Bank Account
func (s *BankService) CreateAccount(userID uint, bankName, accountNumber string, initialBalance float64) (*models.BankAccount, error) {
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, utils.ErrUserNotFound
	}

	if initialBalance < 0 {
		return nil, utils.ErrInvalidAmount
	}

	account := &models.BankAccount{
		UserID:        userID,
		BankName:      bankName,
		AccountNumber: accountNumber,
		Balance:       initialBalance,
		IsActive:      true,
	}

	err = s.repo.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// 4. Get Account Balance (with ownership check)
func (s *BankService) GetBalance(accountNumber string, loggedInUserID uint) (float64, error) {
	account, err := s.repo.GetAccountByNumber(accountNumber)
	if err != nil {
		return 0, utils.ErrAccountNotFound
	}

	if account.UserID != loggedInUserID {
		return 0, utils.ErrUnauthorizedAccountAccess
	}

	return account.Balance, nil
}

// 5. Money Transfer (with ownership check)
func (s *BankService) Transfer(senderAccNum, receiverAccNum string, amount float64, source string, loggedInUserID uint) (*models.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if amount <= 0 {
		return nil, utils.ErrInvalidAmount
	}

	if senderAccNum == receiverAccNum {
		return nil, utils.ErrSameAccountTransfer
	}

	sender, err := s.repo.GetAccountByNumber(senderAccNum)
	if err != nil {
		return nil, utils.ErrAccountNotFound
	}

	// Security: The sender account must belong to the authenticated user!
	if sender.UserID != loggedInUserID {
		return nil, utils.ErrUnauthorizedAccountAccess
	}

	if !sender.IsActive {
		return nil, utils.ErrAccountInactive
	}

	if sender.Balance < amount {
		return nil, utils.ErrInsufficientBalance
	}

	receiver, err := s.repo.GetAccountByNumber(receiverAccNum)
	if err != nil {
		return nil, fmt.Errorf("receiver account: %w", utils.ErrAccountNotFound)
	}

	if !receiver.IsActive {
		return nil, fmt.Errorf("receiver account: %w", utils.ErrAccountInactive)
	}

	txn := &models.Transaction{
		UserID:            sender.UserID,
		SenderAccountID:   sender.ID,
		ReceiverAccountID: receiver.ID,
		ReceiverName:      receiver.BankName + " - " + receiver.AccountNumber,
		Amount:            amount,
		TransactionType:   "TRANSFER",
		Source:            source,
		Status:            "COMPLETED",
	}

	err = s.repo.TransferFunds(sender.ID, receiver.ID, amount, txn)
	if err != nil {
		return nil, utils.ErrTransactionFailed
	}

	return txn, nil
}

// 6. Get Account Statement (with ownership check)
func (s *BankService) GetStatement(accountNumber string, loggedInUserID uint) ([]models.Transaction, error) {
	account, err := s.repo.GetAccountByNumber(accountNumber)
	if err != nil {
		return nil, utils.ErrAccountNotFound
	}

	if account.UserID != loggedInUserID {
		return nil, utils.ErrUnauthorizedAccountAccess
	}

	return s.repo.GetTransactionsByAccount(account.ID)
}
