package main

import (
	"fmt"
	"time"
)

// import "structs"

//struct are custom DS

type order struct{
    id string
    amount float32
    status string
    createdAt time.Time // nanosecond precision + accurate
}

func main()  {
	//group multiple filed
    // myorder := order{
    //     id:"1",
    //     amount:50000,
    //     status: "received",
    // }
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
    
}