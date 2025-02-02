package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
// func processIt(fn func(a int)int){
// 	fn(1)
// }
func processIt() func(a int)int{
	return func(a int)int{
		return 4
	}
}
func main() {
	result := add(3, 5)
	fmt.Println(result)
	// fn:=func(a int) int{
	// 	return 2
	// }
	fn:=processIt()
	fn(7)

}