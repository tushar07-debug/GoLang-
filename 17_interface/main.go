package main

import "fmt"

type payment struct {
}

func (p *payment) makepayment(amount float32) {
	razorpaypayment := razorpay{}
	razorpaypayment.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay",amount)
}
func main() {
	newPayment:=payment{}
	newPayment.makepayment(1000)
}
