package utils

import "errors"

var (
	ErrInsufficientBalance       = errors.New("insufficient balance in account")
	ErrAccountNotFound           = errors.New("bank account not found")
	ErrAccountInactive           = errors.New("bank account is inactive")
	ErrSameAccountTransfer       = errors.New("cannot transfer money to the same account")
	ErrInvalidAmount             = errors.New("transaction amount must be greater than zero")
	ErrUserNotFound              = errors.New("user not found")
	ErrTransactionFailed         = errors.New("transaction processing failed")
	ErrInvalidCredentials        = errors.New("invalid email or password")
	ErrUnauthorizedAccountAccess = errors.New("unauthorized: account does not belong to you")
)
