package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing-task", id)
}
func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i <= 10; i++ {
		waitGroup.Add(1)
		go task(i,&waitGroup)
		// go func (i int){
		// 	fmt.Println(i)
		// }(i)
	}

	waitGroup.Wait()
	// time.Sleep(time.Second * 2) //blocking nhi h unoerdereed ,, concurrently run ho rhi h isloye proepr run nhi ho rha h

}
