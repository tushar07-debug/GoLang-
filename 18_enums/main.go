package main

import "fmt"

//enumerated type
//enums->types+conts

type orderstatus int
const(
	Received orderstatus = iota //predeclared identifire for unsigned integer
	Confirmed
	Prepared
	Declared
	// delivered orderstatus = "delivered"
	// paid orderstatus = "paid"
)
func changeorderstatus(status orderstatus)  {
	fmt.Println("order status changed to",status)	
}
func main(){
	changeorderstatus(Received)
}