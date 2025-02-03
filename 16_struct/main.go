package main

import (
	"fmt"
)

// import "structs"

//struct are custom DS

// type order struct{
//     id string
//     amount float32
//     status string
//     createdAt time.Time // nanosecond precision + accurate
// }

//constructor
// func neworder(id string, amount float32, status string) *order{
//     //initial setup goes here
//     myorder := order{
//         id: id,
//         amount:amount,
//         status: status,
//     }
//     return &myorder
// }

//to connect with the struct order we have to writre (o) o is the frist alphabert of the strcut name i.e order-> o

//reciever type
// func (o *order) changestatus(status string){
//     o.status=status
// }
// func (o *order) getamount() float32{
//     return o.amount
// }

func main() {
	//group multiple filed
	// if u dont set any filed , default value is zero
	//int => 0, float=> 0, string => empty i.e "", bool=> false
	// myorder := order{
	//     id:"1",
	//     amount:50000,
	//     status: "received",
	// }
	// myorder.changestatus("confirmed")
	// fmt.Println(myorder.getamount())
	// fmt.Println(myorder)

	// myorder.createdAt=time.Now()
	// fmt.Println(myorder.status)
	// myorder2:=order{
	//     id:"2",
	//     amount:5000,
	//     status: "delivered",
	//     createdAt: time.Now(),
	// }
	// myorder.status="paid"
	// fmt.Println("order strcut",myorder)
	// fmt.Println("order strcut",myorder2)

	// myorder:=neworder("1",50.00,"received")
	// fmt.Println(myorder.amount)

    //inline struct
	langauge := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(langauge)
}
