package main

import "fmt"

// numbered sequence of specific length
func main() {
	//added zero value by default in int type 
	var nums [4]int

	//aadded false by default in array
	var num [4]bool
	// fmt.Println(len(nums))

	nums[0] = 1
	nums[1]=2
	fmt.Println(nums)
	fmt.Println(num)

	//decalre it in single line
	numed :=[3]int{1,2,3}
	fmt.Println(numed)

	//2d array
	nume:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nume)

	//fixed size, that is predictable
	//memory optimization
	// constant time access
}
