package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("task", id)
}
func main() {
	for i := 0; i <= 10; i++ {
		go task(i)

		// go func (i int){
		// 	fmt.Println(i)
		// }(i)
	}
	time.Sleep(time.Second * 2) //blocking nhi h unoerdereed ,, concurrently run ho rhi h isloye proepr run nhi ho rha h 
}
