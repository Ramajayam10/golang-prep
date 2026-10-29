package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotification struct{
	EmailAddress string
}

type SMSNotification struct{
	PhoneNumber string
}


func(e EmailNotification) Send(message string){
	fmt.Printf("Sending Email to %s: %s\n", e.EmailAddress, message)
}

func(s SMSNotification) Send(message string){
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
}

func NotifyUser(n Notifier, message string){
	n.Send(message)
}

func main(){
	email := EmailNotification{
		EmailAddress: "anushka@test.com",
	}

	sms := SMSNotification{
		PhoneNumber: "1234567890",
	}
	
	NotifyUser(email, "hello")
	NotifyUser(sms, "hello")
	
	var n Notifier = email
	if _, ok := n.(EmailNotification); ok {
		fmt.Println("It is an Email Notification")
	}
}	