package main

import "fmt"

func printslice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
// func printslicestring(items []string)()  {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
func main() {

	nums := []int{1, 2, 3, 4, 5}
	// names:=[]string{"golang","typescript"}
	/// boool "=[]bool{true,false}"
	printslice(nums)
	// printslicestring(names)
}
