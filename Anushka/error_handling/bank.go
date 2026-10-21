package errorhandling

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Insufficient funds for this withdrawal")

func Withdraw(balance int, amount int) (int, error) {
	if amount > balance {
		return balance, ErrInsufficientFunds
	}
	return balance - amount, nil
}


func TransferMoney(senderBalance int, receiverBalance int, amount int) (int, int, error) {
	newSenderBalance, err := Withdraw(senderBalance, amount)
	if err != nil {
		return senderBalance, receiverBalance, fmt.Errorf("transfer failed: %w", err)
	}

	return newSenderBalance, receiverBalance + amount, nil
}
