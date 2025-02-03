package main

import "fmt"

// interface are contract

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// gateway razorpay
	gateway paymenter
}

//open close principle - violits

//concrete implemenetation of razor pay and stripe
func (p *payment) makepayment(amount float32) {
	// razorpaypayment := razorpay{}
	// stripePayment := stripe{}
	// stripePayment.pay(amount)
	// razorpaypayment.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct {
// }

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }
type fakepayment struct {
}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fakepayment")
}

type paypal struct {
}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}
func main() {
	// stripePayment := stripe{}
	// razorpayPayment := razorpay{}
	// fakeGW := fakepayment{}
	paypals:= paypal{}
	newPayment := payment{
		// newPayment.makepayment(1000)
		gateway: paypals,
	}
	newPayment.makepayment(1000)
}
