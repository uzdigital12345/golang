package main

import "fmt"

type MessageService interface {
	SendChargeNotification(int) bool
}

type SMSService struct{}

type MyService struct {
	messageService MessageService
}

func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending Production charge notification")
	return true
}

func (a MyService) ChargeCustomer(value int ) bool {

	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging customer for the value of %d\n", value)

	return true
}

func main()  {
	smsService := SMSService{}

	myService := MyService{smsService}

	myService.ChargeCustomer(100)
}