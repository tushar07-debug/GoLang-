package main

import "fmt"

type payment struct {
}

func (p *payment) makepayment(amount float32) {
	// razorpaypayment := razorpay{}
	stripePayment := stripe{}
	stripePayment.pay(amount)
	// razorpaypayment.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay",amount)
}

type stripe struct{

}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe",amount)
}
func main() {
	newPayment:=payment{}
	newPayment.makepayment(1000)
}
