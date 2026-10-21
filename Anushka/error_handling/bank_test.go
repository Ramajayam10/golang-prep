package errorhandling

import (
	"errors"
	"testing"

)

func TestWithdraw_Success(t *testing.T){
	balance := 100
	amount := 50
	expected := (balance-amount)

	newBalance, err := Withdraw(balance,amount)

	if err != nil{
		t.Errorf("Expected no error , but got %v", err)
	}
	if newBalance != expected{
		t.Errorf("Expected balance %d but got %d",expected,newBalance)
	}
}

func TestWithdraw_InsufficientFunds(t *testing.T){
	balance := 50
	amount := 100

	_,err := Withdraw(balance,amount)

	if err == nil {
		t.Errorf("Expected insufficient funds error , but got no error")
	}
	if !errors.Is(err,ErrInsufficientFunds){
		t.Errorf("Expected error %v but got %v",ErrInsufficientFunds,err)
	}
}


func TestTransferMoney_Failure(t *testing.T) {
	senderBal := 30
	receiverBal := 50
	amount := 100 
	_, _, err := TransferMoney(senderBal, receiverBal, amount)

	if err == nil {
		t.Errorf("Expected transfer error, got nil")
	}

	if !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("Expected error to wrap ErrInsufficientFunds, but got: %v", err)
	}
}
