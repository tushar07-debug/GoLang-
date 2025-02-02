package main

import "fmt"


func sum(nums ...int) int {
	total:=0
	for _, num:=range nums{
		total+=num
	}	
	return total
}
func main() {
	fmt.Println(1,2,3,4,5,"hello") // variadic funtion -> n no of paramneterr can pass
	nums:= []int{3,3,3,3,3}
	result:=sum(nums...)
	fmt.Println(result)
}