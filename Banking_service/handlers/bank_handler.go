package handlers

import (
	"errors"
	"net/http"

	"banking-service/services"
	"banking-service/utils"

	"github.com/gin-gonic/gin"
)

type BankHandler struct {
	service *services.BankService
}

// Constructor
func NewBankHandler(service *services.BankService) *BankHandler {
	return &BankHandler{service: service}
}

// Helper to safely extract logged-in user ID from Gin context
func getUserIDFromContext(c *gin.Context) (uint, error) {
	val, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("unauthorized: missing user ID in context")
	}

	switch v := val.(type) {
	case float64:
		return uint(v), nil
	case uint:
		return v, nil
	case int:
		return uint(v), nil
	default:
		return 0, errors.New("unauthorized: invalid user ID format")
	}
}

// 1. POST /api/v1/users (Registration)
type CreateUserRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Address     string `json:"address"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

func (h *BankHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(req.Name, req.PhoneNumber, req.Email, req.Address, req.State, req.Country, req.Password, req.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

// 2. POST /api/v1/login (Authentication)
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *BankHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// 3. POST /api/v1/accounts (Protected)
type CreateAccountRequest struct {
	BankName       string  `json:"bank_name" binding:"required"`
	AccountNumber  string  `json:"account_number" binding:"required"`
	InitialBalance float64 `json:"initial_balance"`
}

func (h *BankHandler) CreateAccount(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.service.CreateAccount(userID, req.BankName, req.AccountNumber, req.InitialBalance)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account opened successfully",
		"data":    account,
	})
}

// 4. POST /api/v1/transfer (Protected)
type TransferRequest struct {
	SenderAccount   string  `json:"sender_account" binding:"required"`
	ReceiverAccount string  `json:"receiver_account" binding:"required"`
	Amount          float64 `json:"amount" binding:"required,gt=0"`
	Source          string  `json:"source" binding:"required"` // e.g. G-PAY, PHONEPE, ATM
}

func (h *BankHandler) Transfer(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txn, err := h.service.Transfer(req.SenderAccount, req.ReceiverAccount, req.Amount, req.Source, userID)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUnauthorizedAccountAccess):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

		case errors.Is(err, utils.ErrInsufficientBalance),
			errors.Is(err, utils.ErrSameAccountTransfer),
			errors.Is(err, utils.ErrInvalidAmount):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, utils.ErrAccountNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		case errors.Is(err, utils.ErrAccountInactive):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transfer failed"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer completed successfully",
		"data":    txn,
	})
}

// 5. GET /api/v1/accounts/balance?account_number=HDFC101 (Protected)
func (h *BankHandler) GetBalance(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accountNumber := c.Query("account_number")
	if accountNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account_number is required"})
		return
	}

	balance, err := h.service.GetBalance(accountNumber, userID)
	if err != nil {
		if errors.Is(err, utils.ErrUnauthorizedAccountAccess) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, utils.ErrAccountNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_number": accountNumber,
		"balance":        balance,
	})
}

// 6. GET /api/v1/accounts/statement?account_number=HDFC101 (Protected)
func (h *BankHandler) GetStatement(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accountNumber := c.Query("account_number")
	if accountNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account_number is required"})
		return
	}

	txns, err := h.service.GetStatement(accountNumber, userID)
	if err != nil {
		if errors.Is(err, utils.ErrUnauthorizedAccountAccess) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, utils.ErrAccountNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_number": accountNumber,
		"transactions":   txns,
	})
}
