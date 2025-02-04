package main

import (
	"fmt"
)

//send
// func processNum(numChan chan int) {
//     for num:=range numChan{
//         fmt.Println("processing number", num)
//         time.Sleep(time.Second)
//     }
// 	// fmt.Println("processign number", <-numChan)
// }

//recieve
// func sum(result chan int, num1 int, num2 int) {
// 	numresult := num1 + num2
// 	result <- numresult

// }

//goroutine synchonizer
// func task(done chan bool) {
//     defer func(){done<-true}()
//     fmt.Println("processing...")
//     // done<-true
// }

// func emailsender(emailchan chan string, done chan bool) {
// 	defer func() { done <- true }()
//     // <-done
//     // emailchan<-"somehting@gamil.com"
// 	for email := range emailchan {
// 		fmt.Println("email sent to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case num1 := <-chan1:
			fmt.Println(num1)
		case str1 := <-chan2:
			fmt.Println(str1)
		}
	}

	// emailchan := make(chan string, 100) //buffered channel
	// done := make(chan bool)

	// go emailsender(emailchan, done)
	// for i := 0; i < 4; i++ {
	// 	emailchan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")
	// close(emailchan)
	// <-done

	// emailchan <- "hell@gamil.com"
	// email := <-emailchan
	// fmt.Println(email) //non blovking

	// done := make(chan bool)
	// go task(done)
	// <-done//block hojaega yah pr aakr

	// result := make(chan int)
	// go sum(result, 10, 20)
	// result1 := <-result //blocking
	// fmt.Println(result1)

	// messageChan := make(chan string)
	// messageChan <- "hello" //channels are blocking until second side is not ready to recidve it
	// message := <-messageChan
	// fmt.Println(message)

	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 51
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// time.Sleep(time.Second * 2)
}
