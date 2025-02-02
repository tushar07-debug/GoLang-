package main

import "fmt"

//by value-> distinct copy we get
// func change(num int) {
// 	num = 5
// 	fmt.Println("in change", num)
// }
//send num by refernce to chage the number

//by refrence
func changeref(num *int) {
	*num = 5 //derefrnce done  y giving * to numbecause num is pointer
	fmt.Println("in change", *num)

}
func main() {
	num := 1
	changeref(&num)

	// fmt.Println("memory address", &num)
	fmt.Println("after change", num)
}
